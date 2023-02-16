/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/thatstoasty/health-tracker/shared/models"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
)

// exerciseCmd represents the exercise command
var exerciseCmd = &cobra.Command{
	Use:   "exercise",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exercise called")
	},
}

type Exercise struct {
	Name      string   `json:"name"`
	Primary   []string `json:"primary"`
	Secondary []string `json:"secondary"`
	Tertiary  []string `json:"tertiary"`
	Type      string   `json:"type"`
	Variation string   `json:"variation"`
}

type Exercises struct {
	Exercises []Exercise `json:"exercises"`
}

var ExercisesCmd = &cobra.Command{
	Use:   "exercises",
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

		// we initialize our Exercises array
		var exercises Exercises

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		error := json.Unmarshal(byteValue, &exercises)
		if error != nil {
			fmt.Println(error)
		}

		fmt.Println(exercises.Exercises)

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

		ctx := context.Background()

		// we iterate through every user within our users array and
		// print out the user Type, their name, and their facebook url
		// as just an example
		for i := 0; i < len(exercises.Exercises); i++ {
			fmt.Println("\n----")
			fmt.Println(exercises.Exercises[i])
			fmt.Println("Exercise Name: " + exercises.Exercises[i].Name)
			fmt.Printf("Primary: %v", exercises.Exercises[i].Primary)
			fmt.Printf("\nSecondary: %v", exercises.Exercises[i].Secondary)
			fmt.Printf("\nTertiary: %v", exercises.Exercises[i].Tertiary)
			fmt.Printf("\nType Of: %v", exercises.Exercises[i].Type)
			fmt.Printf("\nVariation Of: %v", exercises.Exercises[i].Variation)
		}

		for _, exercise := range exercises.Exercises {
			fmt.Println("\n----")
			fmt.Println(exercise)
			record := models.SubmitExerciseParams{
				Name:      exercise.Name,
				Type:      sql.NullString{String: exercise.Type, Valid: true},
				Variation: sql.NullString{String: exercise.Variation, Valid: true},
			}
			response, err := queries.SubmitExercise(ctx, record)
			log.Println(response)
			if err != nil {
				log.Fatal(err)
			}
			for _, bodyPart := range exercise.Primary {
				entry := models.SubmitExerciseDetailsParams{
					ExerciseName: exercise.Name,
					BodyPart:     bodyPart,
					Level:        "primary",
				}
				response, err := queries.SubmitExerciseDetails(ctx, entry)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(response)
			}
			for _, bodyPart := range exercise.Secondary {
				entry := models.SubmitExerciseDetailsParams{
					ExerciseName: exercise.Name,
					BodyPart:     bodyPart,
					Level:        "secondary",
				}
				response, err := queries.SubmitExerciseDetails(ctx, entry)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(response)
			}
			for _, bodyPart := range exercise.Tertiary {
				entry := models.SubmitExerciseDetailsParams{
					ExerciseName: exercise.Name,
					BodyPart:     bodyPart,
					Level:        "tertiary",
				}
				response, err := queries.SubmitExerciseDetails(ctx, entry)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(response)
			}
		}
	},
}

type BodyPart struct {
	Name         string `json:"name"`
	Region       string `json:"region"`
	UpperOrLower string `json:"upper_or_lower"`
}

type BodyParts struct {
	BodyParts []BodyPart `json:"body_parts"`
}

var BodyPartsCmd = &cobra.Command{
	Use:   "body-parts",
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

		// we initialize our BodyParts array
		var bodyParts BodyParts

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		error := json.Unmarshal(byteValue, &bodyParts)
		if error != nil {
			fmt.Println(error)
		}

		fmt.Println(bodyParts.BodyParts)

		queries, err := utils.GetQueryInterface()
		if err != nil {
			fmt.Printf("Failed to connect to postgres and create a query interface: %v", err)
		}

		ctx := context.Background()

		for _, bodyPart := range bodyParts.BodyParts {
			fmt.Println("\n----")
			fmt.Println(bodyPart)
			entry := models.SubmitBodyPartParams{
				Name:         bodyPart.Name,
				Region:       bodyPart.Region,
				UpperOrLower: bodyPart.UpperOrLower,
			}
			response, err := queries.SubmitBodyPart(ctx, entry)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(response)
		}
	},
}

func init() {
	rootCmd.AddCommand(exerciseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exerciseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exerciseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
