package cmd

import (
	"fmt"
	"path"

	"github.com/rapde/rap/lib/action"
	"github.com/rapde/rap/lib/utils"
	"github.com/spf13/cobra"
)

var dockerComposeFilePath = path.Join(utils.GetWorkDir(), ".rap", "docker-compose.yml")

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start services",
	Long:  `start services specified in rap.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO 处理错误
		out, err := action.ExecDockerCompose("-f", dockerComposeFilePath, "up", "-d")
		fmt.Println(out, err)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
