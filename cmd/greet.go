package cmd

import "github.com/spf13/cobra"

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Print a greeting message",
	Long:  `This command prints a greeting message to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello, welcome to the Cobra CLI!")
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
