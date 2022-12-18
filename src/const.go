package main

// Assembly Syntax
const (
	CHARS_COMMENT = ";"
)

// Dot-Directives
const (
	DD_LABEL = ".label"
	DD_DEF   = ".def"
)

// 4BID-N Instructions
const (
	ASM_BRK  = 0x0 // Halt the program
	ASM_LDAI = 0x1 // Load immediate value to acc
	ASM_LDAM = 0x2 // Load memory value to acc
	ASM_STA  = 0x3 // Store acc to memory

	ASM_INC = 0x4 // Increment/Decrement acc
	ASM_ADD = 0x5 // Add memory value to acc
	//ASM_001 = 0x6 //
	//ASM_SHL = 0x7 //

	ASM_NOT = 0x8 // Bitwise NOT
	ASM_ORA = 0x9 // Bitwise OR memory value and acc
	ASM_AND = 0xA // Bitwise AND memory value and acc
	ASM_SHF = 0xB // Bitwise shift (l/r & rot based on high bits)

	ASM_SLP  = 0xC // Sleeps for B seconds at A scale
	ASM_BNE  = 0xD // Skips B many instructions if acc does not equal A
	ASM_JMPI = 0xE // Jump to immediate program location
	ASM_JMPM = 0xF // Jump to memory jump vector
)

// 4BID-N F-Page Addresses
const (
	FPG_P1_DPAD = 0x0 // Player 1 Direction-Pad
	FPG_P1_BTNS = 0x1 // Player 1 Buttons
	FPG_P2_DPAD = 0x2 // Player 2 Direction-Pad
	FPG_P2_BTNS = 0x3 // Player 2 Buttons

	FPG_SCR_X   = 0x4 // Screen X Coord
	FPG_SCR_Y   = 0x5 // Screen Y Coord
	FPG_SCR_VAL = 0x6 // Screen Pixel Value
	FPG_SCR_OPT = 0x7 // Screen Options

	FPG_BEEP_VOL = 0x8 // Beeper Volume
	FPG_BEEP_PTC = 0x9 // Beeper Pitch
	//FPG_BEEP_ = 0xA // Beeper reserved
	//FPG_BEEP_ = 0xB // Beeper reserved

	FPG_RAND = 0xC // Pseudo-Random Number
	//FPG_ = 0xD //
	//FPG_ = 0xE //
	//FPG_ = 0xF //
)
