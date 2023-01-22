package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sonar/pkg"
	"sonar/register"

	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:     "images [flags]",
	Aliases: []string{"image"},
	Short:   "Search a docker registry for image repositories",
	Long: `Search a docker registry for image repositories. 
If no image is specified, then this will return all images on the repository.
If an image is specified, that image will be returned along with any tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		imageRunner(args)
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

func imageRunner(args []string) {
	var request pkg.RequestBuilder
	client := http.Client{}
	registry := register.GetRegistry()

	if len(args) == 0 {
		request = pkg.GetRequestBuilder("image")
	} else {
		image := args[0]
		registry.Image = image
		request = pkg.GetRequestBuilder("tags")
	}
	director := pkg.NewDirector(request)
	httpRequest := director.BuildRequest()

	r, err := http.NewRequest(httpRequest.Method, registry.Registry+httpRequest.Path, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(r)
	if err != nil {
		log.Fatalln(err)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", string(out))
}
