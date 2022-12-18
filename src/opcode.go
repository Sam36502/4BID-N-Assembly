package main

type ArgType int

type Opcode struct {
	Name   string
	Binary byte
	Arg1   ArgType
	Arg2   ArgType
}

const (
	ARG_NONE      = ArgType(0)
	ARG_IMMEDIATE = ArgType(1)
	ARG_ADDRESS   = ArgType(2)
)

var ALL_OPCODES = []Opcode{
	{
		Name:   "BRK",
		Binary: ASM_BRK,
		Arg1:   ARG_NONE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAM,
		Arg1:   ARG_ADDRESS,
		Arg2:   ARG_ADDRESS,
	},
	{
		Name:   "STA",
		Binary: ASM_STA,
		Arg1:   ARG_ADDRESS,
		Arg2:   ARG_ADDRESS,
	},
	{
		Name:   "INC",
		Binary: ASM_INC,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_IMMEDIATE,
	},
	{
		Name:   "ADD",
		Binary: ASM_ADD,
		Arg1:   ARG_ADDRESS,
		Arg2:   ARG_ADDRESS,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "LDA",
		Binary: ASM_LDAI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
}
