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
		Registry := register.GetRegistry()

		resp, err := http.Get(Registry.Registry + "/v2/_catalog")
		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if len(args) == 0 {
			fmt.Printf("%s", string(body))
		} else {
			image := args[0]
			Registry.Image = image

			resp, err := http.Get(pkg.UrlBuilder("tags"))
			if err != nil {
				log.Fatalln(err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("%s", string(body))
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
