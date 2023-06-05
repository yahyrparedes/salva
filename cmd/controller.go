package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

const (
	TemplateController = "/controller.tmpl"
)

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Create file controller",
	Long:  "Create file controller for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		ValidateExistOrCreateDirectory(PathController)

		var controller = ""

		if len(args) >= 1 && args[0] != "" {
			controller = args[0]
		}

		if len(controller) == 0 {
			fmt.Println("Set name controller")
			return
		}
		name := strings.ToLower(controller)
		runes := []rune(name)
		runes[0] = unicode.ToUpper(runes[0])
		controller = string(runes)
		data := Data{
			ListName:   "Get" + controller + "s",
			DetailName: "Get" + controller,
			UpdateName: "Update" + controller,
			DeleteName: "Delete" + controller,
		}

		//ProcessTemplate(
		//	TemplateController,
		//	PathController+name+".go",
		//	data)
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
