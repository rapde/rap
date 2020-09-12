package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/rapde/rap/lib/config"
	"github.com/rapde/rap/lib/service"
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
	var (
		sc                = bufio.NewScanner(os.Stdin)
		supportedServices = service.Mgr().SupportedServices()
		svcs, _           = toSupportedSvcMap(supportedServices)
		//svcLen            = len(svcs)
	)

	sort.Strings(svcs)
	fmt.Println("Currently supported services:")
	for i := range svcs {
		fmt.Printf("%d. %s\n", i+1, svcs[i])
	}
	fmt.Print("Please input number(e.g. 1) to select service: ")
	if sc.Scan() {
		fmt.Println(sc.Text())
	}
}

func toSupportedSvcMap(supportedSvcs []config.ServiceKey) ([]string, map[string][]string) {
	m := map[string][]string{}
	svcs := []string{}
	for _, v := range supportedSvcs {
		svcs = append(svcs, v.Name())
		tmp := m[v.Name()]
		tmp = append(tmp, v.Version())
		m[v.Name()] = tmp
	}

	return svcs, m
}
