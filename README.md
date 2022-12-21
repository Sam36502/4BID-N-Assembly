# 4BID-N-Assembly
Assembler for the 4BID-N fantasy console

## Binary File Format (`.4bb`)
This way of storing 4BOD/4BID-N instructions is really just storing
the instructions in 2 Bytes; There are 4 bits wasted per instruction,
but for the sake of simplicity and given the limited size of
4BOD binaries it doesn't make sense to pack the data any smaller.
(128 bytes wasted for a full 256-instruction program)

This format is simply a series of 2-Byte instructions as follows:

    Byte 1:
    0000      Top 4 bits are empty
        0000  Bottom 4 bits of first byte stores the instruction opcode

    Byte 2:
    0000      Top 4 bits are argument 1
        0000  Bottom 4 bits are argument 2

This pattern repeats every 2 Bytes for a total of 512 Bytes.
Whether empty instructions after the end of the program is stored
is left unspecified.
