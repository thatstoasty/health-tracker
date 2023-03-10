/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

// getCmd represents the get command
var getTrainingWeight = &cobra.Command{
	Use:   "training-weight",
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

		rirMapping := utils.GetRIRMapping()
		weightTable := utils.GetWeightTable()
		percentage := weightTable[reps-1][rirMapping[rir]]
		trainingWeight := fmt.Sprintf("%.2f", weight*percentage)

		fmt.Printf("Percentage: %v", percentage*100)
		fmt.Printf("\nTraining Weight: %v", trainingWeight)
	},
}

var getTrainingMax = &cobra.Command{
	Use:   "training-max",
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

		rirMapping := utils.GetRIRMapping()
		weightTable := utils.GetWeightTable()
		percentage := weightTable[reps-1][rirMapping[rir]]
		trainingMax := fmt.Sprintf("%.2f", weight/percentage)

		fmt.Printf("Percentage: %v", percentage*100)
		fmt.Printf("\nTraining Max: %v", trainingMax)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getTrainingWeight)
	getCmd.AddCommand(getTrainingMax)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
