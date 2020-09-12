package cmd

import (
	"fmt"

	"github.com/rapde/rap/lib/config"
	"github.com/rapde/rap/lib/service"
	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "prepare rap env",
	Long:  `init config rap.yaml and .rap folder, download related tools and docker images`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("env called")
		fmt.Println(service.Mgr().SupportedServices())
		fmt.Println(service.Mgr().GenDockerComposeConfig(
			[]*config.Service{
				{
					Name:    "alexmysql",
					Service: "mysql",
					Version: "5.7",
				},
			},
		))
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
