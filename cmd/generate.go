/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
/*	For anyone who finds themselves here, wondering what the difference is ... using --ss="one" --ss="two,three" as an example:

StringSlice* - will result in []string{"one", "two", "three"}
StringArray* - will result in []string{"one", "two,three"}
Learn more at:

https://godoc.org/github.com/spf13/pflag#StringArray
https://godoc.org/github.com/spf13/pflag#StringSlice*/

var generateCmd = &cobra.Command{
	Use:   "generate <file> <name>",
	Short: "Genere les fichiers pour slim_bridge",
	Long: `Genere Action, Command, Event, All 
	exemple : generator generate action test --appPath={} --basePath={} --domain={} --eventName={} --params="param1, param2, param3,..."
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v", args) 
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("appPath", "a", "", "")
	rootCmd.PersistentFlags().StringP("basePath", "b", "", "")
	rootCmd.PersistentFlags().StringP("domain", "d", "", "")

	rootCmd.MarkPersistentFlagRequired("appPath")
	rootCmd.MarkPersistentFlagRequired("basePath")
	rootCmd.MarkPersistentFlagRequired("domain")

	rootCmd.AddCommand(generateCmd)
}



