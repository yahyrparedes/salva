package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	TemplateRouterFile = "router.tmpl"
	PathRouter         = "./pkg/routes/"
	Router             = "RouterTemplate"
)

var routeCmd = &cobra.Command{
	Use:   "router",
	Short: "Create file router",
	Long:  "Create file router for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		print()
		input := GetArg(args)
		data := GenerateData(input)
		CreateRouter(input, data)
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)
}

func CreateRouter(name string, data Data) {
	ValidateExistOrCreateDirectory(PathRouter)
	file := PathRouter + name + ".go"
	ProcessTemplateString(Router, file, data)
	//ProcessFileTemplate(TemplateRouterFile, file, data)
	fmt.Printf("Success create router %s\n", name)
}
