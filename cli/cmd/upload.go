/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/thatstoasty/health-tracker/cli/load"
	"github.com/thatstoasty/health-tracker/shared/queries"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var submitCompositionFile = &cobra.Command{
	Use:   "submit-compositions",
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

func init() {
	rootCmd.AddCommand(submitCompositionFile)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
