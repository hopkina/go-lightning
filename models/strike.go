package models

import (
	"time"
)

type LightningStrike struct {
	ID         string    `json:"id" gorm:"primary_key"`
	StrikeTime time.Time `json:"strikeTime"`
	XCoord     float64   `json:"xCoord"`
	YCoord     float64   `json:"yCoord"`
}

type CreateStrikeInput struct {
	StrikeTime time.Time `json:"strikeTime" binding:"required"`
	XCoord     float64   `json:"xCoord" binding:"required"`
	YCoord     float64   `json:"yCoord" binding:"required"`
}

type UpdateStrikeInput struct {
	StrikeTime time.Time `json:"strikeTime"`
	XCoord     float64   `json:"xCoord"`
	YCoord     float64   `json:"yCoord"`
}

type LightningStrikeCountry struct {
	Count   string `json:"count"`
	Country string `json:"country"`
}
