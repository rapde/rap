package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "prepare rap env",
	Long:  `init config rap.yaml and .rap folder, download related tools and docker images`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("env called")
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
