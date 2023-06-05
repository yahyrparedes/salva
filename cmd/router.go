package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

const (
	TemplateRoute = "router.tmpl"
)

var routeCmd = &cobra.Command{
	Use:   "router",
	Short: "Create file router",
	Long:  "Create file router for optimize your development code.",
	Run: func(cmd *cobra.Command, args []string) {
		ValidateExistOrCreateDirectory(PathRouter)

		var input = ""

		if len(args) >= 1 && args[0] != "" {
			input = args[0]
		}

		if len(input) == 0 {
			fmt.Println("Set name controller")
			return
		}
		name := strings.ToLower(input)
		runes := []rune(name)
		runes[0] = unicode.ToUpper(runes[0])
		input = string(runes)
		data := Data{
			ListName:   "Get" + input + "s",
			DetailName: "Get" + input,
			UpdateName: "Update" + input,
			DeleteName: "Delete" + input,
			RouteName:  name,
		}

		ProcessTemplate(
			PathTemplate+TemplateRoute,
			PathRouter+name+".go",
			data)

		fmt.Printf("Success create router name %s\n", name)
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)
}
