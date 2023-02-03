/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/thatstoasty/health-tracker/cli/load"
	"github.com/thatstoasty/health-tracker/shared/queries"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var submitCompositionCmd = &cobra.Command{
	Use:   "submit-composition",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		weight, _ := cmd.Flags().GetString("weight")
		bodyfat, _ := cmd.Flags().GetString("bodyfat")

		entry := queries.SubmitCompositionParams {SubmittedOn: date, Weight: weight, Bodyfat: bodyfat}

		db, err := utils.GetDBConnection()
		if err != nil {
			log.Fatal(err)
		}

		queries := queries.New(db)
		ctx := context.Background()

		composition, err := queries.SubmitComposition(ctx, entry)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(composition)
	},
}

// getCmd represents the get command
var submitNutritionCmd = &cobra.Command{
	Use:   "submit-nutrition",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		calories, _ := cmd.Flags().GetInt16("calories")
		protein, _ := cmd.Flags().GetInt16("protein")
		carbohydrates, _ := cmd.Flags().GetInt16("carbohydrates")
		fats, _ := cmd.Flags().GetInt16("fats")


		entry := queries.SubmitNutritionParams {
			SubmittedOn: date, 
			Calories: calories, 
			Protein: sql.NullInt16 {Int16: protein, Valid: true}, 
			Carbohydrate: sql.NullInt16 {Int16: carbohydrates, Valid: true}, 
			Fat: sql.NullInt16 {Int16: fats, Valid: true},
		}

		db, err := utils.GetDBConnection()
		if err != nil {
			log.Fatal(err)
		}

		queries := queries.New(db)
		ctx := context.Background()

		nutrition, err := queries.SubmitNutrition(ctx, entry)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(nutrition)
	},
}

// getCmd represents the get command
var submitCompositionFileCmd = &cobra.Command{
	Use:   "submit-composition-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		records := load.GetRecordsFromFile(path)
		fmt.Println(records)
		list := load.CreateCompositionList(records)
		fmt.Printf("%+v\n", list)

		db, err := utils.GetDBConnection()
		if err != nil {
			log.Fatal(err)
		}

		queries := queries.New(db)
		ctx := context.Background()

		for _, entry := range list {
			composition, err := queries.SubmitComposition(ctx, entry)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(composition)
		}
	},
}

// getCmd represents the get command
var submitNutritionFileCmd = &cobra.Command{
	Use:   "submit-nutrition-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		records := load.GetRecordsFromFile(path)
		fmt.Println(records)
		list := load.CreateNutritionList(records)
		fmt.Printf("%+v\n", list)

		db, err := utils.GetDBConnection()
		if err != nil {
			log.Fatal(err)
		}

		queries := queries.New(db)
		ctx := context.Background()

		for _, entry := range list {
			nutrition, err := queries.SubmitNutrition(ctx, entry)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(nutrition)
		}
	},
}


func init() {
	rootCmd.AddCommand(submitCompositionCmd)
	rootCmd.AddCommand(submitCompositionFileCmd)
	rootCmd.AddCommand(submitNutritionCmd)
	rootCmd.AddCommand(submitNutritionFileCmd)

	submitCompositionCmd.Flags().String("date", "", "Weigh-in date.")
	submitCompositionCmd.Flags().String("weight", "", "Bodyweight in lbs.")
	submitCompositionCmd.Flags().String("bodyfat", "", "Bodyfat percentage.")

	submitNutritionCmd.Flags().String("date", "", "Date of submission.")
	submitNutritionCmd.Flags().Int16("calories", 0, "Calories consumed on date.")
	submitNutritionCmd.Flags().Int16("protein", 0, "Protein consumed on date.")
	submitNutritionCmd.Flags().Int16("carbohydrates", 0, "Carbohydrates consumed on date.")
	submitNutritionCmd.Flags().Int16("fats", 0, "Fats consumed on date.")


	submitCompositionFileCmd.Flags().String("path", "", "Path to the CSV file containing the composition entries.")
	submitNutritionFileCmd.Flags().String("path", "", "Path to the CSV file containing the nutrition entries.")




	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
