package cmd

import (
	//	"fmt"

	"fmt"

	"github.com/rapde/rap/lib/action"
	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "show all services",
	Long:  `show all services`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO 处理错误
		out, err := action.ExecDockerCompose("-f", dockerComposeFilePath, "ps")
		fmt.Println(out, err)
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
}
