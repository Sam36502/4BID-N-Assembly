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

// Parses a file into a program and returns a list of errors
func ParseFile(filename string) (Program, []error) {
	errors := []error{}
	program := Program{}

	labels := map[string]byte{}
	definitions := DEFAULT_DEFINITIONS

	// Load File
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		errors = append(errors, err)
		return program, errors
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
			}
			labels[fields[1]] = insPointer
			continue

		case DD_DEF:
			if len(fields) != 3 {
				errors = append(errors, FormatSyntaxError("definition dot-directive requires exactly 2 arguments: name & value"))
			}
			definitions[fields[1]] = fields[2]
			continue

		}

		// Parse Arguments
		arg1, isImmediate, err := parseArgument(definitions, fields[1])
		if err != nil {
			errors = append(errors, err)
			continue
		}
		arg2, _, err := parseArgument(definitions, fields[2])
		if err != nil {
			errors = append(errors, err)
			continue
		}

		// Parse Opcode
		var opcode Opcode
		for _, op := range ALL_OPCODES {
			if strings.EqualFold(op.Name, fields[0]) &&
				((isImmediate && op.Arg1 == ARG_IMMEDIATE) || (!isImmediate && op.Arg1 != ARG_IMMEDIATE)) {
				opcode = op
			}
		}

		// Generate instruction
		ins := Instruction{
			Ins:  opcode.Binary,
			Arg1: arg1,
			Arg2: arg2,
		}

		program = append(program, ins)
	}

	return program, errors
}

// Returns value, whether it's immediate or an address
func parseArgument(definitions map[string]string, str string) (byte, bool, error) {
	str = strings.TrimSpace(str)
	originalStr := str

	// Check if it's an immediate value
	var value byte = 0
	isImmediate := false
	if strings.HasPrefix(str, CHARS_IMMEDIATE) {
		isImmediate = true
		str = strings.TrimPrefix(str, CHARS_IMMEDIATE)
	}

	// Check if it's a defined string
	if def, exists := definitions[str]; exists {
		str = def
	}

	// Check if it's a numeric literal
	for base, prefix := range NUM_PREFIXES {
		if strings.HasPrefix(str, prefix) {
			str = strings.TrimPrefix(str, prefix)
			value, err := strconv.ParseUint(str, base, 64)
			if err != nil {
				return 0, isImmediate, FormatSyntaxError(fmt.Sprintf("Invalid numeric literal '%s'", originalStr))
			}
			return byte(value), isImmediate, nil
		}
	}

	return value, isImmediate, FormatSyntaxError(fmt.Sprintf("Invalid argument '%s'", originalStr))
}

func FormatSyntaxError(msg string) error {
	return fmt.Errorf("[%d] Syntax Error: %s", g_lineNr, msg)
}
