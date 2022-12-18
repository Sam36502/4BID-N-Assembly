package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Instruction struct {
	Ins  byte
	Arg1 byte
	Arg2 byte
}

type Program []Instruction

// Parses a file into a program and returns a list of errors
func ParseFile(filename string) (Program, []error) {
	errors := []error{}
	program := Program{}

	labels := map[string]byte{}
	definitions := map[string]string{}

	// Load File
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		errors = append(errors, err)
		return program, errors
	}

	// Parse file line by line
	insPointer := byte(0)
	lines := strings.Split(string(data), "\n")
	for lineNr, line := range lines {

		// Remove comments
		i := strings.Index(line, CHARS_COMMENT)
		if i != -1 {
			line = line[:i]
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)

		// Check for dot-directives
		switch fields[0] {

		case DD_LABEL:
			if len(fields) != 2 {
				errors = append(errors, FormatSyntaxError(lineNr, "label dot-directive requires exactly 1 argument: label name"))
			}
			labels[fields[1]] = insPointer
			continue

		case DD_DEF:
			if len(fields) < 3 {
				errors = append(errors, FormatSyntaxError(lineNr, "definition dot-directive requires at least 2 arguments: name & value"))
			}
			definitions[fields[1]] = strings.Join(fields[2:], " ")
			continue

		}

		// Parse Opcodes and handle arguments

	}

	return program, errors
}

func FormatSyntaxError(lineNr int, msg string) error {
	return fmt.Errorf("[%d] Syntax Error: %s", lineNr, msg)
}
