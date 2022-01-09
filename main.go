package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// lightning_strike represents data about a lightning strike.
type lightningStrike struct {
	ID         string    `json:"id"`
	StrikeTime time.Time `json:"strikeTime"`
	XCoord     float64   `json:"xCoord"`
	YCoord     float64   `json:"yCoord"`
}

// lightning strikes slice to seed lightning strike data.
var lightningStrikes = []lightningStrike{
	{ID: "1", StrikeTime: time.Date(2021, time.November, 10, 23, 55, 0, 0, time.UTC), XCoord: 51.615335, YCoord: -1.193025},
	{ID: "2", StrikeTime: time.Date(2021, time.November, 10, 21, 24, 0, 0, time.UTC), XCoord: 56.099230, YCoord: -3.922570},
	{ID: "3", StrikeTime: time.Date(2021, time.November, 10, 16, 32, 0, 0, time.UTC), XCoord: 52.060492, YCoord: -4.638846},
}

func main() {
	router := gin.Default()
	router.GET("/strikes", getStrikes)
	router.POST("/strikes", postStrikes)

	router.Run("localhost:8080")
}

// getStrikes responds with the list of all lightning strikes as JSON.
func getStrikes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, lightningStrikes)
}

// postStrikes adds a lightning strike from JSON received in the request body.
func postStrikes(c *gin.Context) {
	var newStrike lightningStrike

	// Call BindJSON to bind the received JSON to
	// newStrike.
	if err := c.BindJSON(&newStrike); err != nil {
		return
	}

	// Add the new lightning strike to the slice.
	lightningStrikes = append(lightningStrikes, newStrike)
	c.IndentedJSON(http.StatusCreated, newStrike)
}
