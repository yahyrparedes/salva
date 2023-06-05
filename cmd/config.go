package cmd

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

const (
	PathTemplate   = "/templates/"
	PathController = "./pkg/controllers/"
	PathRouter     = "./pkg/routes/"
)

type Data struct {
	ListName   string
	DetailName string
	UpdateName string
	DeleteName string
	RouteName  string
}

func ProcessTemplate(fileName string, outputFile string, data Data) {
	tmpl := template.Must(template.ParseFiles(fileName))
	f, _ := os.Create(outputFile)
	err := tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
}

func ValidateExistOrCreateDirectory(path string) {
	var err error
	flag.Parse()

	err = os.MkdirAll(path, 0775)
	if err != nil {
		fmt.Printf("create %s error: %s\n", path, err)
		panic(err)
	}
}
