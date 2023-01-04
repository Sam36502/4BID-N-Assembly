# 4BID-N-Assembly
Assembler for the 4BID-N fantasy console

## Downloads
 - #### None Yet...

## Getting Started
To begin using the 4nsm 4BID-N assembler, download it
and, if you like, add its path to your system's PATH
variable (if on linux a system, running `make install`
in the cloned repo should make a symlink to `/usr/bin/4nsm`)

Now here's what a basic assembly program typically looks like:

    ;
    ;   Simple program to draw a diagonal line across the screen
    ;
    
    .label loop
    
            LDA     #1                      ; Set the currently selected pixel on
            STA     _scr_val    _fpage
    
            LDA     _scr_x      _fpage
            BNE     #$F         #1          ; Stop the program if we've reached the edge
              BRK
    
            IDC     #1                      ; Increment current X/Y value
            STA     _scr_x      _fpage      
            STA     _scr_y      _fpage
    
            JMP     #loop                    ; Jump back to top of loop

As you can hopefully see, semicolons (`;`) indicate comments.

The `.label` line is a dot-directive creating the name "loop" which contains
the program line number of the line it's written on. Then, at the bottom, we
do a literal (indicated by the `#`) jump to that line, starting the loop over again.

This program just increments both the X and Y screen variables on the F-page and
writes a 1 to the screen, which draws a diagonal line across the screen.

Then, once it's reached the bottom left corner (15, 15) it halts. I've indented
the branch instructions after the `BNE`, but this isn't necessary. In fact, no
precise tab size/length is required, and they can be tabs or spaces as you prefer.

Once you've made your program, run

    4nsm -o myprogram.4bb myprogram.4nsm

and the assembler will take the assembly in "myprogram.4nsm" and output an assembled
binary to "myprogram.4bb". If you don't specify the output file with `-o`, it defaults
to a.out.

Good luck and have fun, lol.
