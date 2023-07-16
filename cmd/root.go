package cmd

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "salva",
	Short: "A quick and easy start to increase your development speed.",
	Long: `
Try salva and generate files for your API.
A quick and easy start to increase your development speed.
Create Controller, Model, Router, Service, Repository and Testing very easy.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringArrayP("project", "p", nil, "name your project")

}

func GetArg(args []string) string {
	var input = ""

	if len(args) >= 1 && args[0] != "" {
		input = args[0]
	}

	if len(input) == 0 {
		fmt.Println("Please enter name")
		os.Exit(0)
	}
	return strings.ToLower(input)
}

func GenerateData(input string) Data {
	dir := GetDirExecuteCommand()
 
	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])
	name := string(runes)
	data := Data{
		ListName:   "Get" + name + "s",
		DetailName: "Get" + name,
		UpdateName: "Update" + name,
		DeleteName: "Delete" + name,
		Name:       name,
		Path:       input,
		Dir:        dir,
	}
	return data
}
