package commands

import "github.com/spf13/cobra"

var sampCmd = &cobra.Command{
	Use:   "samp [string]",
	Short: "A sample command",

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println("Sample command executed with argument:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(sampCmd)
}
