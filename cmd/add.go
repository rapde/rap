package cmd

import (
	//	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new service",
	Long:  `Interactively create a new service with supported services`,
	Run: func(cmd *cobra.Command, args []string) {
		interactivelyAddNewService()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func interactivelyAddNewService() {

}
