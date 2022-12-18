/*
Copyright © 2022 Bismarck
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	FLAG_DEBUG = "debug"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "4bid-asm",
	Short: "Assembles a .4sm into a binary file .4bb",
	Long: `Takes an assembly file and creates a binary that
can be read by the 4BID-N fantasy console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Called")
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
	rootCmd.Flags().BoolP(FLAG_DEBUG, "d", false, "Prints out the program binary once its assembled")
}
