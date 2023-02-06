// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package queries

import (
	"time"

	"github.com/tabbed/pqtype"
	"github.com/thatstoasty/health-tracker/shared/types"
)

type TrackerComposition struct {
	SubmittedOn  string         `json:"submittedOn"`
	Weight       string         `json:"weight"`
	Bodyfat      string         `json:"bodyfat"`
	Neck         types.NullString `json:"neck"`
	Shoulders    types.NullString `json:"shoulders"`
	LeftBicep    types.NullString `json:"leftBicep"`
	RightBicep   types.NullString `json:"rightBicep"`
	LeftTricep   types.NullString `json:"leftTricep"`
	RightTricep  types.NullString `json:"rightTricep"`
	LeftForearm  types.NullString `json:"leftForearm"`
	RightForearm types.NullString `json:"rightForearm"`
	Chest        types.NullString `json:"chest"`
	Waist        types.NullString `json:"waist"`
	LeftQuad     types.NullString `json:"leftQuad"`
	RightQuad    types.NullString `json:"rightQuad"`
	LeftCalf     types.NullString `json:"leftCalf"`
	RightCalf    types.NullString `json:"rightCalf"`
	CretTs       time.Time      `json:"-"`
	UpdtTs       time.Time      `json:"-"`
}

type TrackerExercise struct {
	Name   string    `json:"name"`
	CretTs time.Time `json:"cretTs"`
	UpdtTs time.Time `json:"updtTs"`
}

type TrackerExerciseDetail struct {
	ExerciseName string    `json:"exerciseName"`
	BodyPart     string    `json:"bodyPart"`
	Level        string    `json:"level"`
	CretTs       time.Time `json:"cretTs"`
	UpdtTs       time.Time `json:"updtTs"`
}

type TrackerExercisePerformed struct {
	ID            int32          `json:"id"`
	SetID         int32          `json:"setID"`
	ExerciseName  string         `json:"exerciseName"`
	Reps          int16          `json:"reps"`
	Weight        int16          `json:"weight"`
	RepsInReserve types.NullString `json:"repsInReserve"`
	CretTs        time.Time      `json:"cretTs"`
	UpdtTs        time.Time      `json:"updtTs"`
}

type TrackerNutrition struct {
	SubmittedOn    string                `json:"submittedOn"`
	Calories       int16                 `json:"calories"`
	Protein        types.NullInt16         `json:"protein"`
	Carbohydrate   types.NullInt16         `json:"carbohydrate"`
	Fat            types.NullInt16         `json:"fat"`
	Micronutrients pqtype.NullRawMessage `json:"micronutrients"`
	CretTs         time.Time             `json:"cretTs"`
	UpdtTs         time.Time             `json:"updtTs"`
}

type TrackerProgram struct {
	Name   string    `json:"name"`
	CretTs time.Time `json:"cretTs"`
	UpdtTs time.Time `json:"updtTs"`
}

type TrackerProgramDetail struct {
	ProgramName string    `json:"programName"`
	WorkoutName string    `json:"workoutName"`
	CretTs      time.Time `json:"cretTs"`
	UpdtTs      time.Time `json:"updtTs"`
}

type TrackerSetPerformed struct {
	ID        int32     `json:"id"`
	WorkoutID int32     `json:"workoutID"`
	GroupID   int16     `json:"groupID"`
	CretTs    time.Time `json:"cretTs"`
	UpdtTs    time.Time `json:"updtTs"`
}

type TrackerWorkout struct {
	Name        string    `json:"name"`
	ProgramName string    `json:"programName"`
	CretTs      time.Time `json:"cretTs"`
	UpdtTs      time.Time `json:"updtTs"`
}

type TrackerWorkoutDetail struct {
	WorkoutName  string    `json:"workoutName"`
	GroupID      int16     `json:"groupID"`
	ExerciseName string    `json:"exerciseName"`
	Sets         int16     `json:"sets"`
	Reps         int16     `json:"reps"`
	Weight       int16     `json:"weight"`
	CretTs       time.Time `json:"cretTs"`
	UpdtTs       time.Time `json:"updtTs"`
}

type TrackerWorkoutPerformed struct {
	ID          int32     `json:"id"`
	SubmittedOn time.Time `json:"submittedOn"`
	WorkoutName string    `json:"workoutName"`
	CretTs      time.Time `json:"cretTs"`
	UpdtTs      time.Time `json:"updtTs"`
}
