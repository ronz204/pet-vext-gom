package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vext",
	Short: "vext short example",
	Long:  "vext long example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from vext command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command", err)
		os.Exit(1)
	}
}
