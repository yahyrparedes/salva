package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	TemplateRepositoryFile = "repository.tmpl"
	PathRepository         = "./pkg/repositories/"
	Repository             = "RepositoryTemplate"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Create file repository",
	Long:  "Create file repository for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		input := GetArg(args)
		data := GenerateData(input)
		CreateRepository(input, data)
	},
}

func init() {
	rootCmd.AddCommand(repositoryCmd)
}

func CreateRepository(name string, data Data) {
	ValidateExistOrCreateDirectory(PathRepository)
	file := PathRepository + name + ".go"
	ProcessTemplateString(Repository, file, data)
	//ProcessFileTemplate(TemplateRepositoryFile, file, data)
	fmt.Printf("Success create repository %s\n", name)
}
