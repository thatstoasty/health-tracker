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
	"github.com/thatstoasty/health-tracker/shared/models"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
)

// nutritionCmd represents the nutrition command
var nutritionCmd = &cobra.Command{
	Use:   "nutrition",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nutrition called")
	},
}

// getCmd represents the get command
var NutritionFileCmd = &cobra.Command{
	Use:   "nutrition-entries-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		records := load.GetRecordsFromFile(path)
		list := load.CreateNutritionList(records)

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

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

// getCmd represents the get command
var NutritionCmd = &cobra.Command{
	Use:   "nutrition-entry",
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

		entry := models.SubmitNutritionParams{
			SubmittedOn:  date,
			Calories:     calories,
			Protein:      sql.NullInt16{Int16: protein, Valid: true},
			Carbohydrate: sql.NullInt16{Int16: carbohydrates, Valid: true},
			Fat:          sql.NullInt16{Int16: fats, Valid: true},
		}

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

		ctx := context.Background()

		nutrition, err := queries.SubmitNutrition(ctx, entry)
		if err != nil {
			fmt.Printf("Failed to submit a nutrition entry: %v", err)
		}
		log.Printf("Nutrition entry submitted successfully! %v", nutrition)
	},
}



func init() {
	rootCmd.AddCommand(nutritionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nutritionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nutritionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
