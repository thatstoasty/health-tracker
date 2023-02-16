/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/thatstoasty/health-tracker/shared/models"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// programCmd represents the program command
var programCmd = &cobra.Command{
	Use:   "program",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("program called")
	},
}

var ProgramCmd = &cobra.Command{
	Use:   "program",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		// Open our jsonFile
		jsonFile, err := os.Open(path)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully opened the file.")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// read our opened jsonFile as a byte array.
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
		}

		parsed := gjson.ParseBytes(byteValue)
		json := parsed.Raw
		programName := gjson.Get(json, "name").Str

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

		ctx := context.Background()

		workouts := gjson.Get(json, "workouts")
		fmt.Printf("Inserting program: %s\n", programName)

		response, err := queries.SubmitProgram(ctx, programName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response)

		workouts.ForEach(func(key, value gjson.Result) bool {
			raw := value.Raw
			workout := models.SubmitWorkoutParams{
				Name:        gjson.Get(raw, "name").Str,
				ProgramName: programName,
			}
			log.Println(workout)

			log.Println("Submitting Workout")
			_, err := queries.SubmitWorkout(ctx, workout)
			if err != nil {
				log.Fatal(err)
			}

			programWorkoutLink := models.SubmitProgramDetailsParams{
				ProgramName: programName,
				WorkoutName: gjson.Get(raw, "name").Str,
			}

			_, err2 := queries.SubmitProgramDetails(ctx, programWorkoutLink)
			if err2 != nil {
				log.Println("Submitting program details")
				log.Fatal(err2)
			}

			exercises := gjson.Get(raw, "exercises")
			exercises.ForEach(func(key, value gjson.Result) bool {
				exercise := value.Raw
				details := models.SubmitWorkoutDetailsParams{
					WorkoutName:  gjson.Get(raw, "name").Str,
					GroupID:      int16(gjson.Get(exercise, "group_id").Num),
					ExerciseName: gjson.Get(exercise, "name").Str,
					Sets:         int16(gjson.Get(exercise, "sets").Num),
					Reps:         int16(gjson.Get(exercise, "reps").Num),
					Weight:       sql.NullInt16{Int16: int16(gjson.Get(exercise, "weight").Num), Valid: true},
				}
				log.Println(details)
				_, err := queries.SubmitWorkoutDetails(ctx, details)
				if err != nil {
					log.Fatal(err)
				}
				return true
			})

			return true // keep iterating
		})

	},
}

func init() {
	rootCmd.AddCommand(programCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// programCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// programCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
