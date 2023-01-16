package types

import (
	"github.com/shopspring/decimal"
)


// Composition 
type Composition struct {
	Date         string 	     `json:"date"`
	Weight       decimal.Decimal `json:"weight"`
	Bodyfat      decimal.Decimal `json:"bodyfat"`
	Neck         decimal.Decimal `json:"neck,omitempty"`
	Shoulders    decimal.Decimal `json:"shoulders,omitempty"`
	LeftBicep    decimal.Decimal `json:"leftBicep,omitempty"`
	RightBicep   decimal.Decimal `json:"rightBicep,omitempty"`
	LeftTricep   decimal.Decimal `json:"leftTricep,omitempty"`
	RightTricep  decimal.Decimal `json:"rightTricep,omitempty"`
	LeftForearm  decimal.Decimal `json:"leftForearm,omitempty"`
	RightForearm decimal.Decimal `json:"rightForearm,omitempty"`
	Chest        decimal.Decimal `json:"chest,omitempty"`
	Waist        decimal.Decimal `json:"waist,omitempty"`
	LeftQuad     decimal.Decimal `json:"leftQuad,omitempty"`
	RightQuad    decimal.Decimal `json:"rightQuad,omitempty"`
	LeftCalf     decimal.Decimal `json:"leftCalf,omitempty"`
	RightCalf    decimal.Decimal `json:"rightCalf,omitempty"`
}

// Nutrition 
type Nutrition struct {

}

// Exercise 
type Exercise struct {

}

// Workout 
type Workout struct {

}
