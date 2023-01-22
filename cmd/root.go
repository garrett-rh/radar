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
	// Runs as a pre-req to every action using sonar
	// Reads from the global register to check if the -k or --insecure flag was passed in
	// If it was, skips TLS cert verification for all requests.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		registry := register.GetRegistry()
		if registry.TlsNoVerify {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
	},
	// Calls rootRunner and passes args through
	// All logic is in a separate function for ease of testing
	Run: func(cmd *cobra.Command, args []string) {
		rootRunner()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used to hold &  set flags.
// Also sets global vars in the registry
func init() {
	Registry := register.GetRegistry()
	rootCmd.PersistentFlags().BoolVarP(&Registry.TlsNoVerify, "insecure", "k", false, "Set this flag to true to decline verification on TLS certs (default false)")
	rootCmd.PersistentFlags().StringVarP(&Registry.Registry, "registry", "r", "", "Describes the location of the docker registry")
	if err := rootCmd.MarkPersistentFlagRequired("registry"); err != nil {
		log.Fatalln(err)
	}
}

// Runs a connectivity check against the docker registry.
// Anything outside of a 200 is no good
func rootRunner() {
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
		log.Fatalln("Registry returned status code 401 Unauthorized. Login Required")
	}
}
