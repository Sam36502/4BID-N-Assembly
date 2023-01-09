package asm

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	Ins  byte
	Arg1 byte
	Arg2 byte
}

type Program []Instruction

var g_lineNr = 0
var g_definitions = DEFAULT_DEFINITIONS
var g_labels = map[string]byte{}
var g_missingLabels = map[string][]byte{}
var g_branches = []byte{}
var g_returnVecs = map[string]byte{}
var g_currSub = ""

// Parses a file into a program and returns a list of warnings and a list of errors
func ParseFile(filename string) (Program, []string, []error) {
	errors := []error{}
	warns := []string{}
	program := Program{}

	// Load File
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		errors = append(errors, err)
		return program, warns, errors
	}

	insPointer := byte(0)
	lines := strings.Split(string(data), "\n")

	// Scan file for labels
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, DD_LABEL) || strings.HasPrefix(line, DD_SUB) {
			i := strings.Index(line, CHARS_COMMENT)
			if i != -1 {
				line = line[:i]
			}
			line = strings.TrimSpace(line)
			fields := strings.Fields(line)
			if len(fields) != 2 {
				errors = append(errors, fmt.Errorf("[pre-parse label-scan]: label/sub dot-directive requires exactly 1 argument: label name"))
				continue
			}
			g_missingLabels[fields[1]] = []byte{}
		}
	}

	// Parse file line by line
	var line string
	for g_lineNr, line = range lines {

		// Remove comments
		i := strings.Index(line, CHARS_COMMENT)
		if i != -1 {
			line = line[:i]
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		// Check for dot-directives
		switch fields[0] {

		case DD_SUB:
			if len(fields) != 2 {
				errors = append(errors, fmt.Errorf("subroutine dot-directive requires exactly 1 argument: label name"))
				continue
			}
			g_returnVecs[fields[1]] = byte(len(g_returnVecs))
			g_currSub = fields[1]
			fallthrough

		case DD_LABEL:
			if len(fields) != 2 {
				errors = append(errors, fmt.Errorf("label dot-directive requires exactly 1 argument: label name"))
				continue
			}
			g_labels[fields[1]] = insPointer

			// Update all instructions that didn't know the correct insPointer before
			// (i.e. all JMPs and such that tried to use this label before it was declared)
			if _, exists := g_missingLabels[fields[1]]; exists {
				for _, n := range g_missingLabels[fields[1]] {
					program[n].Arg1 = insPointer % 16
					program[n].Arg2 = (insPointer >> 4) % 16
				}
				delete(g_missingLabels, fields[1])
			}

			continue

		case DD_DEF:
			if len(fields) != 3 {
				errors = append(errors, FormatSyntaxError("definition dot-directive requires exactly 2 arguments: name & value"))
				continue
			}
			if def, ex := g_definitions[fields[1]]; ex {
				warns = append(warns, FormatWarning(fmt.Sprintf("definition '%s' redefined", def)))
				continue
			}
			g_definitions[fields[1]] = fields[2]
			continue

		case DD_EBR:
			if len(g_branches) == 0 {
				errors = append(errors, FormatSyntaxError("end-branch dot-directive used, but no empty BNE's found"))
				continue
			}
			branchLine := g_branches[len(g_branches)-1]
			g_branches = g_branches[:len(g_branches)-1]
			skipLines := insPointer - branchLine - 1
			program[branchLine].Arg2 = skipLines
			continue

		case DD_RTS:
			if g_currSub == "" {
				errors = append(errors, FormatSyntaxError("return-from-subroutine (.rts) called outside of a subroutine"))
				continue
			}
			returnVector := g_returnVecs[g_currSub]
			program = append(program, Instruction{
				Ins:  ASM_JMPM,
				Arg1: (returnVector << 1) | 0,
				Arg2: (returnVector << 1) | 1,
			})
			insPointer++
			g_currSub = ""
			continue

		case DD_JSR:
			if len(fields) != 2 {
				errors = append(errors, FormatSyntaxError("jump-to-subroutine dot-directive (.jsr) requires exactly 1 argument: subroutine name"))
				continue
			}
			srLine, exists := g_labels[fields[1]]
			if !exists {
				if _, exists := g_missingLabels[fields[1]]; exists {
					g_missingLabels[fields[1]] = append(g_missingLabels[fields[1]], insPointer+4)
				} else {
					errors = append(errors, FormatSyntaxError(fmt.Sprintf("No subroutine (label) named '%s' found to jump to", fields[1])))
					continue
				}
			}

			returnVector := insPointer + 5
			retVecAddr := g_returnVecs[fields[1]]
			program = append(program, []Instruction{
				{
					Ins:  ASM_LDAI,
					Arg1: returnVector % 16,
				},
				{
					Ins:  ASM_STA,
					Arg1: (retVecAddr << 1) | 0,
					Arg2: ZERO_PAGE,
				},
				{
					Ins:  ASM_LDAI,
					Arg1: (returnVector >> 4) % 16,
				},
				{
					Ins:  ASM_STA,
					Arg1: (retVecAddr << 1) | 1,
					Arg2: ZERO_PAGE,
				},
				{
					Ins:  ASM_JMPI,
					Arg1: srLine % 16,
					Arg2: (srLine >> 4) % 16,
				},
			}...)

			insPointer += 5
			continue

		}

		// Parse Arguments
		var isImmediate bool = false
		var arg1, arg2 byte = 0, 0
		if len(fields) > 1 {
			str := fields[1]
			if strings.HasPrefix(str, CHARS_IMMEDIATE) {
				isImmediate = true
				str = strings.TrimPrefix(str, CHARS_IMMEDIATE)
			}

			// Check if it's a label
			if insNr, exists := g_labels[str]; exists {
				arg1 = insNr % 16
				arg2 = (insNr >> 4) % 16
			} else if _, exists := g_missingLabels[str]; exists {
				g_missingLabels[str] = append(g_missingLabels[str], insPointer)
			} else {
				arg1, err = parseArgument(str)
				if err != nil {
					errors = append(errors, err)
					continue
				}
			}
		}

		if len(fields) > 2 && arg2 == 0 {
			str := fields[2]
			str = strings.TrimPrefix(str, CHARS_IMMEDIATE)
			arg2, err = parseArgument(str)
			if err != nil {
				errors = append(errors, err)
				continue
			}
		}

		// Parse Opcode
		var opcode Opcode
		for _, op := range ALL_OPCODES {
			if strings.EqualFold(op.Name, fields[0]) &&
				((isImmediate && op.Arg1 == ARG_IMMEDIATE) || (!isImmediate && op.Arg1 != ARG_IMMEDIATE)) {
				opcode = op
			}
		}
		if opcode == (Opcode{}) {
			errors = append(errors, FormatSyntaxError(fmt.Sprintf("No opcode '%s' found", fields[0])))
			continue
		}

		// Generate instruction
		ins := Instruction{
			Ins:  opcode.Binary,
			Arg1: arg1,
			Arg2: arg2,
		}

		// Check for special cases
		if opcode.Binary == ASM_BNE && arg2 == 0 {
			g_branches = append(g_branches, insPointer)
		}

		program = append(program, ins)
		insPointer++
	}

	// Check for unclosed branches
	for _, l := range g_branches {
		errors = append(errors, FormatSyntaxError(fmt.Sprintf("Empty BNE started at instruction %d and never ended.", l)))
	}

	// Check for undeclared labels
	for _, l := range g_missingLabels {
		errors = append(errors, FormatSyntaxError(fmt.Sprintf("label expected by instruction(s) %v that was never declared.", l)))
	}

	return program, warns, errors
}

func parseArgument(str string) (byte, error) {
	str = strings.TrimSpace(str)
	origString := str

	// Check if it's a defined string
	if def, exists := g_definitions[str]; exists {
		str = def
	}

	// Check if it's a numeric literal in non-decimal
	for base, prefix := range NUM_PREFIXES {
		if strings.HasPrefix(str, prefix) {
			str = strings.TrimPrefix(str, prefix)
			value, err := strconv.ParseUint(str, base, 64)
			if err != nil {
				return 0, FormatSyntaxError(fmt.Sprintf("Invalid numeric literal '%s' --> '%s'", origString, str))
			}
			return byte(value), nil
		}
	}

	// ...otherwise, it's a decimal number
	value64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, FormatSyntaxError(fmt.Sprintf("Invalid numeric literal '%s'", str))
	}
	return byte(value64), nil
}

func FormatSyntaxError(msg string) error {
	return fmt.Errorf("[%03d] Syntax Error: %s", g_lineNr+1, msg)
}

func FormatWarning(msg string) string {
	return fmt.Sprintf("[%03d] Warning: %s", g_lineNr+1, msg)
}
