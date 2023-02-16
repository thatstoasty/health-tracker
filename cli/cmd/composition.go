/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"fmt"

	"github.com/thatstoasty/health-tracker/cli/load"
	"github.com/thatstoasty/health-tracker/shared/models"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
)

// compositionCmd represents the composition command
var compositionCmd = &cobra.Command{
	Use:   "composition",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("composition called")
	},
}

// getCmd represents the get command
var CompositionFileCmd = &cobra.Command{
	Use:   "composition-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		records := load.GetRecordsFromFile(path)
		list := load.CreateCompositionList(records)

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

		ctx := context.Background()

		for _, entry := range list {
			_, err := queries.SubmitComposition(ctx, entry)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("Submitted all composition entries successfully!")

	},
}

var CompositionCmd = &cobra.Command{
	Use:   "composition",
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

		entry := models.SubmitCompositionParams{SubmittedOn: date, Weight: weight, Bodyfat: bodyfat}

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

		ctx := context.Background()
		composition, err := queries.SubmitComposition(ctx, entry)
		if err != nil {
			fmt.Printf("Failed to submit a composition entry: %v", err)
		}
		fmt.Printf("Composition entry submitted successfully! %v", composition)
	},
}

func init() {
	rootCmd.AddCommand(compositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compositionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compositionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
