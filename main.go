package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
)

var specTemplate = `openapi: 3.0.1
info:
  title: Generated API
  description: "{{ generateDescription }}"
  version: "0.1"
paths:
{{ $a := index .Methods 0 }}{{ $b := index .Methods 1 }}{{range $i, $r := .Resources}}  {{.}}: {{ range $a }}
    {{ .Name }}:
      tags:
        - "{{ guessTagFromPath $r }}"
      summary: "{{ .Humanized }} {{ guessCollectionResourceFromPath $r }}"
      responses: {{ range $code, $description := .Responses }}
        {{ $code }}:
          description: "{{ $description }}"{{end}}{{end}}
  {{.}}/{id}: {{ range $b }}
    {{ .Name }}:
      tags:
        - "{{ guessTagFromPath $r }}"
      summary: "{{ .Humanized }} {{ guessSingularResourceFromPath $r }}"
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type : string
      responses: {{ range $code, $description := .Responses }}
        {{ $code }}:
          description: "{{ $description }}"{{end}}{{end}}
{{end}}`

type Data struct {
	Resources []string
	Methods   [][]Method
}

type Method struct {
	Name      string
	Humanized string
	Responses map[string]string
}

var unimplementedResponse = map[string]string{
	"501": "Not implemented",
}

var defaultMethods = [][]Method{
	[]Method{
		Method{"get", "List", unimplementedResponse},
		Method{"post", "Add to", unimplementedResponse},
	},
	[]Method{
		Method{"get", "Retrieve a specific", unimplementedResponse},
		Method{"put", "Update a specific", unimplementedResponse},
		Method{"delete", "Delete a specific", unimplementedResponse},
	},
}

func main() {
	d := Data{
		Resources: os.Args[1:],
		Methods:   defaultMethods,
	}
	p := pluralize.NewClient()
	tmpl := template.Must(template.New("generate").Funcs(template.FuncMap{
		"guessTagFromPath": func(path string) string {
			return strings.TrimSpace(strings.Replace(path[0:strings.LastIndex(path, "/")], "/", " ", -1))
		},
		"guessCollectionResourceFromPath": func(path string) string {
			return p.Plural(path[strings.LastIndex(path, "/")+1:])
		},
		"guessSingularResourceFromPath": func(path string) string {
			return p.Singular(path[strings.LastIndex(path, "/")+1:])
		},
		"generateDescription": func() string {
			return "Created by '" + strings.Join(os.Args, " ") + "'"
		},
	}).Parse(specTemplate))
	err := tmpl.Execute(os.Stdout, d)
	if err != nil {
		log.Println(err)
		log.Fatal("Unsupported inputs")
	}
}
