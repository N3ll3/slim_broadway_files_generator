/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"slimfilesgenerator/files"
	"strings"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all <name>",
	Short: "Generer tous les fichiers",
	Long: `example:					
	generator generate command test --appPath={} --basePath={} --domain={} --params="param1, param2, param3" --eventName={}
	nb : echapper les slash et antislash`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")
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

		eventName, err := cmd.Flags().GetString("eventName")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cmd.Flags().GetString("domain"))

		params, err := cmd.Flags().GetStringSlice("params")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(cmd.Flags().GetStringSlice("params"))

			
	var paramsStr string
	paramsStr = ""
	if (len(params) > 0) {
		for i, param := range params {
			params[i] = "\t$"+strings.Trim(param, " ")
		}
		paramsStr = strings.Join(params,"-\t")
	} 

		config := files.Args{
			AppPath: app_path,
			BasePath : base_path,
			Domain : domain,
			Name  : name,
			Params : paramsStr,
			EventName : eventName,
		}

		files.GenerateActionFiles(config)
		files.GenerateCommandFiles(config)
		files.GenerateEventFiles(config)
	},
}

func init() {
	allCmd.PersistentFlags().StringSliceP("params", "p", nil, "")
	allCmd.PersistentFlags().StringP("eventName", "e", "", "")
	generateCmd.AddCommand(allCmd)
}
