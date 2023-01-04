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
var g_branches = []byte{}

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

	// Parse file line by line
	insPointer := byte(0)
	lines := strings.Split(string(data), "\n")
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

		case DD_LABEL:
			if len(fields) != 2 {
				errors = append(errors, FormatSyntaxError("label dot-directive requires exactly 1 argument: label name"))
				continue
			}
			if insNr, ex := g_definitions[fields[1]]; ex {
				errors = append(errors, FormatSyntaxError(fmt.Sprintf("label already defined for instruction nr. 0x%02X", insNr)))
				continue
			}
			g_definitions[fields[1]] = fmt.Sprint(insPointer)
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

		}

		// Parse Arguments
		var isImmediate bool
		var arg1, arg2 byte = 0, 0
		if len(fields) > 1 {
			arg1, isImmediate, err = parseArgument(fields[1])
			if err != nil {
				errors = append(errors, err)
				continue
			}
		}

		if len(fields) > 2 {
			arg2, _, err = parseArgument(fields[2])
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

	return program, warns, errors
}

// Returns value, whether it's immediate or an address
func parseArgument(str string) (byte, bool, error) {
	str = strings.TrimSpace(str)
	origString := str

	// Check if it's an immediate value
	isImmediate := false
	if strings.HasPrefix(str, CHARS_IMMEDIATE) {
		isImmediate = true
		str = strings.TrimPrefix(str, CHARS_IMMEDIATE)
	}

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
				return 0, isImmediate, FormatSyntaxError(fmt.Sprintf("Invalid numeric literal '%s' --> '%s'", origString, str))
			}
			return byte(value), isImmediate, nil
		}
	}

	// ...otherwise, it's a decimal number
	value64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, isImmediate, FormatSyntaxError(fmt.Sprintf("Invalid numeric literal '%s'", str))
	}
	return byte(value64), isImmediate, nil
}

func FormatSyntaxError(msg string) error {
	return fmt.Errorf("[%03d] Syntax Error: %s", g_lineNr+1, msg)
}

func FormatWarning(msg string) string {
	return fmt.Sprintf("[%03d] Warning: %s", g_lineNr+1, msg)
}
