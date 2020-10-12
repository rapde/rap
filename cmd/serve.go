package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/rapde/rap/website"
)

const (
	defaultAddr = "localhost:8000" // Using port :8000 by default
)

// http server address
var httpAddr string

// serveCmd represents the start command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a webserver",
	Long:  `start a webserver to manage rap`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(website.RAR) == 0 {
			log.Fatalln("Oops! The website is not built, please first execute `go generate` in the website directory")
		}

		engine := website.New()
		engine.Run(httpAddr)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringVar(&httpAddr, "http", defaultAddr, "HTTP service address (e.g. \"127.0.0.1:8000\" or just \":8000\")")
}
