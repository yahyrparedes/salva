package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

const (
	TemplateService = "/service.tmpl"
)

var serviceCmd = &cobra.Command{
	Use:   "controller",
	Short: "Create file controller",
	Long:  "Create file controller for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		ValidateExistOrCreateDirectory(PathController)

		var input = ""
		var project = ""

		if len(args) >= 1 && args[0] != "" {
			input = args[0]
		}

		if len(args) >= 2 && args[0] != "" {
			input = args[0]
		}

		if len(input) == 0 {
			fmt.Println("Set name controller")
			return
		}
		path := strings.ToLower(input)
		runes := []rune(path)
		runes[0] = unicode.ToUpper(runes[0])
		name := string(runes)
		data := Data{
			Project:    name,
			ListName:   "Get" + name + "s",
			DetailName: "Get" + name,
			UpdateName: "Update" + name,
			DeleteName: "Delete" + name,
			Name:       name,
			Path:       path,
		}

		ProcessTemplateString(
			ControllerTemplate,
			PathController+name+".go",
			data)
		fmt.Printf("Success create controller name %s\n", name)
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)
}

func CreateService(name string, path string, data Data) {
	ValidateExistOrCreateDirectory(PathService)
	ProcessTemplateString(
		ControllerTemplate,
		PathController+name+".go",
		data)
	fmt.Printf("Success create model name %s\n", name)
}
