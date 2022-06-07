/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"slimfilesgenerator/files"

	"github.com/spf13/cobra"
)

// actionCmd represents the action command
var actionCmd= &cobra.Command{
	Use:   "action <name>",
	Short: "Generer le fichier Action",
	Long: `
	example:					
	generator generate action test --appPath={} --basePath={} --domain={}
	nb : echapper les slash et antislash
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate action called")
		name := args[0]
		app_path, err := cmd.Flags().GetString("appPath")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cmd.Flags().GetString("appPath"))

		base_path, err := cmd.Flags().GetString("basePath")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cmd.Flags().GetString("basePath"))

		domain, err := cmd.Flags().GetString("domain")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(cmd.Flags().GetString("domain"))

		config := files.Args{
			AppPath: app_path,
			BasePath : base_path,
			Domain : domain,
			Name  : name,
			Params  : "",
		}

		files.GenerateActionFiles(config)
	},
}


func init() {
	generateCmd.AddCommand(actionCmd)
}
