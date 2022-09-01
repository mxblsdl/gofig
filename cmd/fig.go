/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

// figCmd represents the fig command
var figCmd = &cobra.Command{
	Use:   "fig",
	Short: "Generate figlet ascii",
	Long:  `Leverage the figlet go package to create figlet ascii word art`,
	Run: func(cmd *cobra.Command, args []string) {
		// This should be a function
		makeFig("hello world")
	},
}

func init() {
	rootCmd.AddCommand(figCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// figCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// figCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func makeFig(s string) {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render(s)
	fmt.Println(renderStr)
}

type Ascii struct {
	word string
}
