package cmd

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

const (
	PathTemplate   = "/cmd/templates/"
	PathController = "./pkg/controllers/"
	PathModel      = "./pkg/models/"
	PathRouter     = "./pkg/routes/"
	PathService    = "./pkg/service/"
	Controller     = "ControllerTemplate"
	Router         = "RouterTemplate"
	Model          = "ModelTemplate"
)

type Data struct {
	Project    string
	ListName   string
	DetailName string
	UpdateName string
	DeleteName string
	Name       string
	Path       string
}

func ProcessTemplate(fileName string, outputFile string, data Data) {
	filePrefix, _ := filepath.Abs(PathTemplate)
	tmpl := template.Must(template.ParseFiles(filePrefix + fileName))
	f, _ := os.Create(outputFile)
	err := tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
}

func ProcessTemplateString(format string, outputFile string, data Data) {
	var tmpl = ""
	if format == Controller {
		tmpl = ControllerTemplate

	}
	if format == Router {
		tmpl = RouterTemplate
	}
	if format == Model {
		tmpl = ModelTemplate
	}

	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	f, _ := os.Create(outputFile)
	err = t.Execute(f, data)
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

const ModelTemplate = `/*
Generate by salva 
*/
package models

import (
	"gorm.io/gorm"
)

type {{.Name}} struct {
	gorm.Model
	
}
`

const RouterTemplate = `/*
Generate by salva 
*/
package routes

import (
	"github.com/gofiber/fiber/v2"
	"vehicular-back-go/pkg/controllers"
)

func {{.Name}}Route(router fiber.Router) {
	router.Get("/{{.Path}}s", controllers.{{.ListName }})
	router.Get("/{{.Path}}/:id", controllers.{{.DetailName}})
	router.Get("/{{.Path}}/:id", controllers.{{.UpdateName}})
	router.Get("/{{.Path}}/:id", controllers.{{.DeleteName}})
}
`
const ControllerTemplate = `/*
Generate by salva
*/
package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func {{.ListName }}(c *fiber.Ctx) error {
	return nil
}

func {{.DetailName}}(c *fiber.Ctx) error {
	return nil
}

func {{.UpdateName}}(c *fiber.Ctx) error {
	return nil
}

func {{.DeleteName}}(c *fiber.Ctx) error {
	return nil
}
	`
