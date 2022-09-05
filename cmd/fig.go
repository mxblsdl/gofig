/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

// figCmd represents the fig command
var figCmd = &cobra.Command{
	Use:   "render",
	Short: "Generate figlet ascii",
	Long:  `Leverage the figlet go package to create figlet ascii word art`,
	Run: func(cmd *cobra.Command, args []string) {
		// This should be a function
		var sb strings.Builder

		for i := 0; i < len(args); i++ {
			sb.WriteString(args[i])
			sb.WriteString(" ")
		}
		makeFig(sb.String())
	},
}

var figInput string

func init() {
	rootCmd.AddCommand(figCmd)
	figCmd.Flags().StringVarP(&figInput, "word", "w", "", "Word to convert")
	// figCmd.MarkFlagRequired("word")
}

func makeFig(s ...string) {
	ascii := figlet4go.NewAsciiRender()
	for _, n := range s {
		renderStr, _ := ascii.Render(n)
		fmt.Println(renderStr)
	}
}

// type Ascii struct {
// 	word string
// }
