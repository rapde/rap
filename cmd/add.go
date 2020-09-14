package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/rapde/rap/lib/config"
	"github.com/rapde/rap/lib/service"
	"github.com/rapde/rap/lib/utils"
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
		supportedServices   = service.Mgr().SupportedServices()
		svcs, svcVersionMap = toSupportedSvcMap(supportedServices)
	)

	sort.Strings(svcs)
	svc := getUserInputSevice(svcs)
	ver := getUserInputVerion(svc, svcVersionMap[svc])
	svcVer := config.NewServiceKey(svc, ver)
	log.Println("Your choice is:", svcVer)

	service := config.Service{
		Service: svc,
		Version: ver,
	}

	if checkWantExtraConfig() {
		cnf := &config.ServiceConfig{}
		cnf.Volumes = getExtraSlice("Volumes")
		cnf.Ports = getExtraSlice("Ports")
		cnf.Environment = getExtraMap("Environment")
		service.Config = cnf
	}

	name := getName()
	if len(name) == 0 {
		name = fmt.Sprintf("%s_%s", svc, utils.GenRandomString(6))
		log.Println("Generate default name:", name)
	}

	service.Name = name
	config.AppConfig.Services = append(config.AppConfig.Services, &service)
	config.AppConfig.Save()
	log.Println("Add service successfully, you can start service now")
}

func getExtraMap(notice string) map[string]string {
	slice := getExtraSlice(notice + " key,value separate by ':'")
	m := map[string]string{}
	for _, v := range slice {
		splits := strings.Split(v, ":")
		if len(splits) != 2 {
			log.Fatal("Invalid Key,Value", v)
		}

		m[splits[0]] = strings.TrimPrefix(splits[1], " ")
	}

	return m
}

func getExtraSlice(notice string) []string {
	slice := []string{}
	sc := bufio.NewScanner(os.Stdin)

	fmt.Printf("Extra %s(Press enter complete input):\n", notice)
	for {
		fmt.Print("- ")
		if !sc.Scan() {
			log.Fatalf("Get extra %s failed %+v", notice, sc.Err())
		}

		vol := sc.Text()
		if len(vol) == 0 {
			break
		}

		slice = append(slice, vol)
	}

	return slice
}

func getName() string {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Please input service name(6-30 characters recommended): ")
	if !sc.Scan() {
		log.Fatal("Get name failed", sc.Err())
	}

	name := sc.Text()
	return name
}

func checkWantExtraConfig() bool {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Do you want extra config?(Y/n): ")
	if !sc.Scan() {
		log.Fatal("check failed", sc.Err())
	}

	input := sc.Text()
	if strings.ToLower(input) == "y" {
		return true
	}

	return false
}

func getUserInputVerion(svc string, versions []string) string {
	sc := bufio.NewScanner(os.Stdin)
selectVersion:
	fmt.Printf("Currently %s supported verions:\n", svc)
	for i, v := range versions {
		fmt.Printf("  %d. %s\n", i+1, v)
	}

	fmt.Print("Please input number(e.g. 1) to select version: ")
	if !sc.Scan() {
		log.Fatal("Select version failed", sc.Err())
	}

	verSelect := sc.Text()
	if len(verSelect) == 0 {
		fmt.Println("You have to select a version first!!")
		goto selectVersion
	}

	idx := getStringNum(verSelect, len(versions))

	return versions[idx-1]
}

func getUserInputSevice(svcs []string) (svc string) {
	sc := bufio.NewScanner(os.Stdin)
selectService:
	fmt.Println("Currently supported services:")
	for i := range svcs {
		fmt.Printf("  %d. %s\n", i+1, svcs[i])
	}

	fmt.Print("Please input number(e.g. 1) to select service: ")
	if !sc.Scan() {
		log.Fatal("Select service failed", sc.Err())
	}

	svcSelect := sc.Text()
	if len(svcSelect) == 0 {
		fmt.Println("You have to select a service first!!")
		goto selectService
	}

	idx := getStringNum(svcSelect, len(svcs))

	svc = svcs[idx-1]
	return
}

func getStringNum(s string, maxLen int) int {
	idx, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Invalid number, convert to digital failed", err)
	}

	if idx < 1 || idx > maxLen {
		log.Fatal("Invalid number, index overflow")
	}

	return idx
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
