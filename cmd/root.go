package cmd

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"sonar/pkg"
	"sonar/register"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "sonar --registry [Registry URL] [command]",
	Short:   "Sonar is a tool used to navigate docker registries",
	Long:    `sonar is an easy way to enumerate a docker registry without direct knowledge of the API.`,
	Example: "sonar --registry https://localhost -k images",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		registry := register.GetRegistry()
		if registry.TlsNoVerify {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		rootRunner(args)
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

func rootRunner(args []string) {
	registry := register.GetRegistry()
	request := pkg.GetRequestBuilder("base")
	director := pkg.NewDirector(request)
	baseRequest := director.BuildRequest()

	client := http.Client{}
	r, err := http.NewRequest(baseRequest.Method, registry.Registry+baseRequest.Path, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(r)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == http.StatusOK {
		log.Println("Connection established!")
	}
	if resp.StatusCode == http.StatusUnauthorized {
		log.Println("Registry returned status code 401 Unauthorized. Login Required")
	}
}
