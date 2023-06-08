package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	TemplateModelFile = "model.tmpl"
	PathModel         = "./pkg/models/"
	Model             = "ModelTemplate"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create file model",
	Long:  "Create file model for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		input := GetArg(args)
		data := GenerateData(input)
		CreateModel(input, data)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}

func CreateModel(name string, data Data) {
	ValidateExistOrCreateDirectory(PathModel)
	file := PathModel + name + ".go"
	ProcessTemplateString(Model, file, data)
	//ProcessFileTemplate(TemplateModelFile, file, data)
	fmt.Printf("Success create model %s\n", name)
}
