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

	"github.com/thatstoasty/health-tracker/cli/load"
	"github.com/thatstoasty/health-tracker/shared/models"
	"github.com/thatstoasty/health-tracker/shared/utils"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
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

// getCmd represents the get command
var compositionEntriesCmd = &cobra.Command{
	Use:   "composition-entries",
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

// getCmd represents the get command
var nutritionEntriesCmd = &cobra.Command{
	Use:   "nutrition-entries",
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

var exercisesCmd = &cobra.Command{
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

var bodyPartsCmd = &cobra.Command{
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

var programCmd = &cobra.Command{
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

// getCmd represents the get command
var compositionEntryCmd = &cobra.Command{
	Use:   "composition-entry",
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

// getCmd represents the get command
var nutritionEntryCmd = &cobra.Command{
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
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.AddCommand(compositionEntriesCmd)
	uploadCmd.AddCommand(nutritionEntriesCmd)
	uploadCmd.AddCommand(exercisesCmd)
	uploadCmd.AddCommand(bodyPartsCmd)
	uploadCmd.AddCommand(programCmd)

	compositionEntriesCmd.Flags().String("path", "", "Path to the CSV file containing the composition entries.")
	nutritionEntriesCmd.Flags().String("path", "", "Path to the CSV file containing the nutrition entries.")
	exercisesCmd.Flags().String("path", "", "Path to the JSON file containing exercise definitions.")
	bodyPartsCmd.Flags().String("path", "", "Path to the JSON file containing body part definitions.")
	programCmd.Flags().String("path", "", "Path to the JSON file containing a program definition.")

	uploadCmd.AddCommand(compositionEntryCmd)
	uploadCmd.AddCommand(nutritionEntryCmd)

	compositionEntryCmd.Flags().String("date", "", "Weigh-in date.")
	compositionEntryCmd.Flags().String("weight", "", "Bodyweight in lbs.")
	compositionEntryCmd.Flags().String("bodyfat", "", "Bodyfat percentage.")

	nutritionEntryCmd.Flags().String("date", "", "Date of submission.")
	nutritionEntryCmd.Flags().Int16("calories", 0, "Calories consumed on date.")
	nutritionEntryCmd.Flags().Int16("protein", 0, "Protein consumed on date.")
	nutritionEntryCmd.Flags().Int16("carbohydrates", 0, "Carbohydrates consumed on date.")
	nutritionEntryCmd.Flags().Int16("fats", 0, "Fats consumed on date.")
}
