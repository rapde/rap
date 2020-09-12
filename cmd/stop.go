package cmd

import (
	"fmt"

	"github.com/rapde/rap/lib/action"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop services",
	Long:  `stop services specified in rap.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO 处理错误
		out, err := action.ExecDockerCompose("-f", dockerComposeFilePath, "down")
		fmt.Println(out, err)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
