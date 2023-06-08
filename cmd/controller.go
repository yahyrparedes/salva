package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	TemplateControllerFile = "controller.tmpl"
	PathController         = "./pkg/controllers/"
	Controller             = "ControllerTemplate"
)

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Generate file controller",
	Long:  "Generate file controller for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		input := GetArg(args)
		data := GenerateData(input)
		CreateController(input, data)
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)
}

func CreateController(name string, data Data) {
	ValidateExistOrCreateDirectory(PathController)
	file := PathController + name + ".go"
	ProcessTemplateString(Controller, file, data)
	//ProcessFileTemplate(TemplateControllerFile, file, data)
	fmt.Printf("Success create controller %s\n", name)
}
