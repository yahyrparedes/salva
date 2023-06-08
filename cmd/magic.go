package cmd

import (
	"github.com/spf13/cobra"
)

var (
	FlagModel      bool
	FlagController bool
	FlagRouter     bool
	FlagService    bool
	FlagRepository bool
)

var magicCmd = &cobra.Command{
	Use:   "magic",
	Short: "Create Controller, Model, Router, Service, Repository",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		input := GetArg(args)
		data := GenerateData(input)

		if FlagModel == false && FlagRouter == false && FlagController == false && FlagService == false && FlagRepository == false {
			CreateModel(input, data)
			CreateRouter(input, data)
			CreateController(input, data)
			CreateService(input, data)
			CreateRepository(input, data)
		}

		if FlagModel {
			CreateModel(input, data)
		}

		if FlagRouter {
			CreateRouter(input, data)
		}

		if FlagController {
			CreateController(input, data)
		}

		if FlagService {
			CreateService(input, data)
		}

		if FlagRepository {
			CreateRepository(input, data)
		}

	},
}

func init() {
	rootCmd.AddCommand(magicCmd)
	magicCmd.Flags().BoolVarP(&FlagModel, "model", "m", false, "Create model")
	magicCmd.Flags().BoolVarP(&FlagController, "controller", "c", false, "Create controller")
	magicCmd.Flags().BoolVarP(&FlagRouter, "router", "r", false, "Create router")
	magicCmd.Flags().BoolVarP(&FlagService, "service", "s", false, "Create service")
	magicCmd.Flags().BoolVarP(&FlagRepository, "repository", "p", false, "Create repository")

}

func Create(name string, path string, data Data) {

}
