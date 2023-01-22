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
	// Calls imageRunner and passes args through
	// All logic is in a separate function for ease of testing
	Run: func(cmd *cobra.Command, args []string) {
		imageRunner(args)
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

// imageRunner takes in the args passed in from the parent command `image(s)`
// If an argument (an image name) is passed in, then then specific info on that image is pulled
// If no arguments are passed in, then the entire image catalog is pulled down from the registry
// Currently, all args after arg[0] are ignored
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
