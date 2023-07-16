package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	TemplateMapperFile = "mapper.tmpl"
	PathMapper         = "./pkg/mappers/"
	Mapper             = "MapperTemplate"
)

var mapperCmd = &cobra.Command{
	Use:   "mapper",
	Short: "Create file mapper",
	Long:  "Create file mapper for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		input := GetArg(args)
		data := GenerateData(input)
		CreateMapper(input, data)
	},
}

func init() {
	rootCmd.AddCommand(mapperCmd)
}

func CreateMapper(name string, data Data) {
	ValidateExistOrCreateDirectory(PathMapper)
	file := PathMapper + name + ".go"
	ProcessTemplateString(Mapper, file, data)
	//ProcessFileTemplate(TemplateMapperFile, file, data)
	fmt.Printf("Success create mapper %s\n", name)
}
