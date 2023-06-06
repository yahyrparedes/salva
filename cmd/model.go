package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

const (
	TemplateModel = "/model.tmpl"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create file model",
	Long:  "Create file model for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {

		var input = ""

		if len(args) >= 1 && args[0] != "" {
			input = args[0]
		}

		if len(input) == 0 {
			fmt.Println("Set name model")
			return
		}
		path := strings.ToLower(input)
		runes := []rune(path)
		runes[0] = unicode.ToUpper(runes[0])
		name := string(runes)
		data := Data{
			ListName:   "Get" + name + "s",
			DetailName: "Get" + name,
			UpdateName: "Update" + name,
			DeleteName: "Delete" + name,
			Name:       name,
			Path:       path,
		}

		CreateModel(name, data)

	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}

func CreateModel(name string, data Data) {
	ValidateExistOrCreateDirectory(PathModel)
	ProcessTemplateString(
		ModelTemplate,
		PathModel+name+".go",
		data)
	fmt.Printf("Success create model name %s\n", name)
}
