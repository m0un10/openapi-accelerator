package main

import (
	"text/template"
	"log"
	"os"
	"strings"
	pluralize "github.com/gertd/go-pluralize"
)

type Data struct {
	Resources []string
	Methods [][]Method
}

type Method struct {
	Name string
	Humanized string
	Responses map[string]string
}

var unimplementedResponse = map[string]string{
	"501":"Not implemented",
}

var defaultMethods = [][]Method{
	[]Method{
		Method{"get","List",unimplementedResponse},
		Method{"post","Add to",unimplementedResponse},
	},
	[]Method{
		Method{"get","Retrieve a specific",unimplementedResponse},
		Method{"put","Update a specific",unimplementedResponse},
		Method{"delete","Delete a specific",unimplementedResponse},
	},
}

func main(){	
	templateFile := "template.yml"
	d := Data{
		Resources: os.Args[1:],
		Methods: defaultMethods,
	}
	p := pluralize.NewClient()
	tmpl := template.New(templateFile).Funcs(template.FuncMap{
		"guessTagFromPath": func(path string) string {
		  return strings.TrimSpace(strings.Replace(path[0:strings.LastIndex(path,"/")],"/"," ",-1))
		},
		"guessCollectionResourceFromPath": func(path string) string {
			return p.Plural(path[strings.LastIndex(path,"/")+1:])
		},
		"guessSingularResourceFromPath": func(path string) string {
			return p.Singular(path[strings.LastIndex(path,"/")+1:])
		},
	})
	tmpl, err := tmpl.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Invalid template: "+templateFile)
	}
	err = tmpl.Execute(os.Stdout, d)
	if err != nil {
		log.Println(err)
		log.Fatal("Unsupported inputs")
	}
}