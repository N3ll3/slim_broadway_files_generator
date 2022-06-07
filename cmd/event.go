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

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event <name>",
	Short: "Generer le fichier event",
	Long: `example:					
	generator generate event test --appPath={} --basePath={} --domain={} --params="param1, param2, param3" --eventName={}
	nb : echapper les slash et antislash`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("event called")
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

		files.GenerateEventFiles(config)
	},
}

func init() {
	eventCmd.PersistentFlags().StringSliceP("params", "p", nil, "")
	eventCmd.PersistentFlags().StringP("eventName", "e", "", "")
	generateCmd.AddCommand(eventCmd)
}
