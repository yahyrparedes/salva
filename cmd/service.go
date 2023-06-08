package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	TemplateServiceFile = "service.tmpl"
	PathService         = "./pkg/services/"
	Service             = "ServiceTemplate"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create file service",
	Long:  "Create file service for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		input := GetArg(args)
		data := GenerateData(input)
		CreateService(input, data)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}

func CreateService(name string, data Data) {
	ValidateExistOrCreateDirectory(PathService)
	file := PathService + name + ".go"
	ProcessTemplateString(Service, file, data)
	//ProcessFileTemplate(TemplateServiceFile, file, data)
	fmt.Printf("Success create service %s\n", name)
}
