package cmd

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	PathTemplate = "templates/*.tmpl"
)

// var templates = template.Must(template.New("T").ParseGlob(PathTemplate))
var templates = template.New("")

type Data struct {
	Project    string
	ListName   string
	DetailName string
	UpdateName string
	DeleteName string
	Name       string
	Path       string
	Dir        string
}

func ProcessFileTemplate(fileName string, outputFile string, data Data) {
	file, _ := os.Create(outputFile)
	err := templates.ExecuteTemplate(file, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
}

func getTemplateString(file string) string {
	var tmpl = ""

	if file == Controller {
		tmpl = TemplateControllerString
	}

	if file == Router {
		tmpl = TemplateRouterString
	}

	if file == Model {
		tmpl = TemplateModelString
	}

	if file == Service {
		tmpl = TemplateServiceString
	}

	if file == Repository {
		tmpl = TemplateRepositoryString
	}

	return tmpl
}

func ProcessTemplateString(format string, outputFile string, data Data) {
	var tmpl = getTemplateString(format)
	t := template.Must(template.New(format).Parse(tmpl))
	f, _ := os.Create(outputFile)
	err := t.Execute(f, data)
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

const TemplateServiceString = `// Generate by salva 
package services

import (
	"{{.Dir}}/pkg/mappers"
	Query{{.Name}} "{{.Dir}}/pkg/repositories"
)

func Get{{.Name}}s() []mappers.{{.Name}} {
	{{.Path}}s := Query{{.Name}}.FindAll()
	m{{.Name}}s := mappers.Request{{.Name}}s({{.Path}}s)

	return m{{.Name}}s
}
`

const TemplateRepositoryString = `// Generate by salva  
package repositories

import (
	"{{.Dir}}/cmd/database"
	"{{.Dir}}/pkg/models"
)

func FindAll() []models.{{.Name}} {
	var {{.Path}}s []models.{{.Name}}
	database.Connection.Find(&{{.Path}}s)
	return {{.Path}}s
}

func FindByID(id int) models.{{.Name}} {
	var {{.Path}} models.{{.Name}}
	database.Connection.Find(&{{.Path}}, "id", id)
	return {{.Path}}
}

`

const TemplateModelString = `// Generate by salva 
package models

import (
	"gorm.io/gorm"
)

type {{.Name}} struct {
	gorm.Model

}
`

const TemplateRouterString = `// Generate by salva
package routes

import (
	"github.com/gofiber/fiber/v2"
	"{{.Dir}}/pkg/controllers"
)

func {{.Name}}Route(router fiber.Router) {
	router.Get("/{{.Path}}s", controllers.{{.ListName }})
	router.Get("/{{.Path}}/:id", controllers.{{.DetailName}})
	router.Get("/{{.Path}}/:id", controllers.{{.UpdateName}})
	router.Get("/{{.Path}}/:id", controllers.{{.DeleteName}})
}
`

const TemplateControllerString = `// Generate by salva
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

func GetPwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	p := strings.Split(path, "/")
	dir := p[len(p)-1]
	return dir
}
func GetPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	p := strings.Split(path, "/")
	dir := p[len(p)-1]
	return dir
}

func GetDirExecuteCommand() string {
	dir := GetPwd()
	if len(dir) != 0 {
		return dir
	}
	dir = GetPath()
	if len(dir) != 0 {
		return dir
	}
	return ""
}
