package asm

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
		Name:   "---",
		Binary: 0b0000,
		Arg1:   ARG_NONE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "---",
		Binary: 0b0000,
		Arg1:   ARG_NONE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "NOT",
		Binary: ASM_NOT,
		Arg1:   ARG_NONE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "ORA",
		Binary: ASM_ORA,
		Arg1:   ARG_ADDRESS,
		Arg2:   ARG_ADDRESS,
	},
	{
		Name:   "AND",
		Binary: ASM_AND,
		Arg1:   ARG_ADDRESS,
		Arg2:   ARG_ADDRESS,
	},
	{
		Name:   "SHF",
		Binary: ASM_SHF,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_NONE,
	},
	{
		Name:   "SLP",
		Binary: ASM_SLP,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_IMMEDIATE,
	},
	{
		Name:   "BNE",
		Binary: ASM_BNE,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_IMMEDIATE,
	},
	{
		Name:   "JMP",
		Binary: ASM_JMPI,
		Arg1:   ARG_IMMEDIATE,
		Arg2:   ARG_IMMEDIATE,
	},
	{
		Name:   "JMP",
		Binary: ASM_JMPM,
		Arg1:   ARG_ADDRESS,
		Arg2:   ARG_ADDRESS,
	},
}
