/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
	},
}


func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.AddCommand(CompositionFileCmd)
	uploadCmd.AddCommand(NutritionFileCmd)
	uploadCmd.AddCommand(ExercisesCmd)
	uploadCmd.AddCommand(BodyPartsCmd)
	uploadCmd.AddCommand(ProgramCmd)
	uploadCmd.AddCommand(WorkoutPerformedCmd)

	CompositionFileCmd.Flags().String("path", "", "Path to the CSV file containing the composition entries.")
	NutritionFileCmd.Flags().String("path", "", "Path to the CSV file containing the nutrition entries.")
	ExercisesCmd.Flags().String("path", "", "Path to the JSON file containing exercise definitions.")
	BodyPartsCmd.Flags().String("path", "", "Path to the JSON file containing body part definitions.")
	ProgramCmd.Flags().String("path", "", "Path to the JSON file containing a program definition.")
	WorkoutPerformedCmd.Flags().String("path", "", "Path to the JSON file containing a program definition.")

	uploadCmd.AddCommand(CompositionFileCmd)
	uploadCmd.AddCommand(NutritionFileCmd)

	CompositionCmd.Flags().String("date", "", "Weigh-in date.")
	CompositionCmd.Flags().String("weight", "", "Bodyweight in lbs.")
	CompositionCmd.Flags().String("bodyfat", "", "Bodyfat percentage.")

	NutritionCmd.Flags().String("date", "", "Date of submission.")
	NutritionCmd.Flags().Int16("calories", 0, "Calories consumed on date.")
	NutritionCmd.Flags().Int16("protein", 0, "Protein consumed on date.")
	NutritionCmd.Flags().Int16("carbohydrates", 0, "Carbohydrates consumed on date.")
	NutritionCmd.Flags().Int16("fats", 0, "Fats consumed on date.")
}
