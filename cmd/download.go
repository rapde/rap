package cmd

import (
	"github.com/rapde/rap/lib/action"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download & build images",
	Long:  `download & build images specified in docker-compose file`,
	Run: func(cmd *cobra.Command, args []string) {
		// action.Ensuredownload()
		action.ExecDockerCompose("-f", dockerComposeFilePath, "pull")
		action.ExecDockerCompose("-f", dockerComposeFilePath, "build")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
