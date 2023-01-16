package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type repos struct {
	Repository []string `json:"repositories"`
}

var imageCmd = &cobra.Command{
	Use:     "images [flags]",
	Aliases: []string{"image"},
	Short:   "Search a docker registry for image repositories",
	Long: `Search a docker registry for image repositories. 
If no image is specified, then this will return all images on the repository.
If an image is specified, that image will be returned along with any tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(Registry + "/v2/_catalog")
		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if len(args) == 0 {
			fmt.Println(string(body))
		} else {
			image := args[0]
			var decoded repos
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &decoded); err != nil {
				log.Fatalln(err)
			}
			resp, err := http.Get(Registry + "/v2/" + image + "/tags/list")
			if err != nil {
				log.Fatalln(err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(string(body))
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
