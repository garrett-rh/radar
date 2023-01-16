package cmd

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	TlsNoVerify bool
	Registry    string
)
var rootCmd = &cobra.Command{
	Use:   "rummage --registry [Registry URL] [command]",
	Short: "Rummage is a tool used to navigate docker registries",
	Long:  `Rummage is an easy way to enumerate a docker registry without direct knowledge of the API.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if TlsNoVerify {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := http.Get(Registry + "/v2")
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
	rootCmd.PersistentFlags().BoolVarP(&TlsNoVerify, "insecure", "k", false, "Set this flag to true to decline verification on TLS certs (default false)")
	rootCmd.PersistentFlags().StringVarP(&Registry, "registry", "r", "", "Describes the location of the docker registry")
	if err := rootCmd.MarkPersistentFlagRequired("registry"); err != nil {
		log.Fatalln(err)
	}
}
