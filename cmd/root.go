package cmd

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"sonar/register"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "sonar --registry [Registry URL] [command]",
	Short:   "Sonar is a tool used to navigate docker registries",
	Long:    `sonar is an easy way to enumerate a docker registry without direct knowledge of the API.`,
	Example: "sonar --registry https://localhost -k images",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		Registry := register.GetRegistry()
		if Registry.TlsNoVerify {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		Registry := register.GetRegistry()

		resp, err := http.Get(Registry.Registry + "/v2")
		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode == http.StatusOK {
			log.Println("Connection established!")
		}
		if resp.StatusCode == http.StatusUnauthorized {
			log.Println("Registry returned status code 401 Unauthorized. Login Required")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Registry := register.GetRegistry()
	rootCmd.PersistentFlags().BoolVarP(&Registry.TlsNoVerify, "insecure", "k", false, "Set this flag to true to decline verification on TLS certs (default false)")
	rootCmd.PersistentFlags().StringVarP(&Registry.Registry, "registry", "r", "", "Describes the location of the docker registry")
	if err := rootCmd.MarkPersistentFlagRequired("registry"); err != nil {
		log.Fatalln(err)
	}
}
