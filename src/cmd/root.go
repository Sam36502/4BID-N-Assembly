/*
Copyright © 2022 Bismarck
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sam36502/4BID-N-Assembly/src/asm"
	"github.com/spf13/cobra"
)

const (
	FLAG_DEBUG  = "debug"
	FLAG_OUTPUT = "outfile"
	FILE_MODE   = 0650
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "4bid-asm",
	Short: "Assembles a .4sm into a binary file .4bb",
	Long: `Takes an assembly file and creates a binary that
can be read by the 4BID-N fantasy console.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get arguments
		if len(args) != 1 {
			fmt.Println("Exactly one argument is required: input assembly file")
			return
		}
		infile := args[0]

		outfile, err := cmd.Flags().GetString(FLAG_OUTPUT)
		if err != nil {
			fmt.Println("Failed to read outfile flag: ", err)
			return
		}

		// Read and parse infile
		prog, warns, errs := asm.ParseFile(infile)
		if len(errs) > 0 {
			fmt.Println("Assembly failed:")
			for _, err := range errs {
				fmt.Printf("  %v\n", err)
			}
			return
		}
		if len(warns) > 0 {
			fmt.Println("Warnings:")
			for _, warn := range warns {
				fmt.Printf("  %s\n", warn)
			}
		}

		// Generate and save output file
		data := make([]byte, len(prog)*2)
		for progi, ins := range prog {
			i := progi * 2
			data[i] = ins.Ins % 16
			data[i+1] = (ins.Arg1%16)<<4 | (ins.Arg2 % 16)
		}

		err = ioutil.WriteFile(outfile, data, FILE_MODE)
		if err != nil {
			fmt.Println("Failed to write outfile: ", err)
			return
		}

		fmt.Printf("Program assembled to '%s'\n", outfile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP(FLAG_DEBUG, "d", false, "Prints out the program binary once its assembled (WIP)")
	rootCmd.Flags().StringP(FLAG_OUTPUT, "o", "a.out", "The file to output to")
}
