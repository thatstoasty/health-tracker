/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thatstoasty/health-tracker/tools/cli/reference"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get-training-weight",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		repsInput := args[0]
		rirInput := args[1]
		weightInput := args[2]

		reps, err := strconv.Atoi(repsInput)
		if err != nil {
			panic(err)
		}
		rir, err := strconv.ParseFloat(rirInput, 64)
		if err != nil {
			panic(err)
		}
		weight, err := strconv.ParseFloat(weightInput, 64)
		if err != nil {
			panic(err)
		}

		rirMapping := reference.GetRIRMapping()
		weightTable := reference.GetWeightTable()

		fmt.Println(weight * weightTable[reps-1][rirMapping[rir]])
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
