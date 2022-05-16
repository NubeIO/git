package cmd

import (
	"os"

	"github.com/fatih/color"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Magenta(err.Error())
		os.Exit(1)
	}
}
