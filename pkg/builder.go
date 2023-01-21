package pkg

import (
	"bytes"
	"log"
	"sonar/register"
	"text/template"
)

func UrlBuilder(endpoint string) string {
	registry := register.GetRegistry()
	var uri bytes.Buffer

	_, err := uri.WriteString(registry.Registry)
	if err != nil {
		log.Fatalln(err)
	}

	if endpoint == "tags" {
		uri.WriteString(getTags(registry.Image))
	}

	return uri.String()
}

func getTags(image string) string {
	var temp bytes.Buffer
	tmpl, err := template.New("url").Parse(register.GetEndpoints().Tags)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(&temp, map[string]interface{}{"Image": string(image)})
	if err != nil {
		log.Fatalln(err)
	}

	return temp.String()
}
