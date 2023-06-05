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
	PathTemplate       = "/cmd/templates/"
	PathController     = "./pkg/controllers/"
	PathRouter         = "./pkg/routes/"
	ControllerTemplate = "ControllerTemplate"
	RouterTemplate     = "RouterTemplate"
)

type Data struct {
	ListName   string
	DetailName string
	UpdateName string
	DeleteName string
	RouteName  string
	Route      string
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
	if format == ControllerTemplate {
		tmpl = `/*
Generate by is_salva @yahyrparedes
*/
package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func {{.ListName }}(c *fiber.Ctx) error {

}

func {{.DetailName}}(c *fiber.Ctx) error {

}

func {{.UpdateName}}(c *fiber.Ctx) error {

}

func {{.DeleteName}}(c *fiber.Ctx) error {


}
	`

	}
	if format == RouterTemplate {
		tmpl = `/*
Generate by is_salva @yahyrparedes
*/
package routes

import (
	"github.com/gofiber/fiber/v2"
	"vehicular-back-go/pkg/controllers"
)

func {{.Route}}Routes(router fiber.Router) {
	router.Get("/{{.RouteName}}s", controllers.{{.ListName }})
	router.Get("/{{.RouteName}}/:id", controllers.{{.DetailName}})
	router.Get("/{{.RouteName}}/:id", controllers.{{.UpdateName}})
	router.Get("/{{.RouteName}}/:id", controllers.{{.DeleteName}})
}
`
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
