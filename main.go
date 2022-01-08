package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// lightning_strike represents data about a lightning strike.
type lightning_strike struct {
	ID          string    `json:"id"`
	Strike_Time time.Time `json:"strikeTime"`
	X_Coord     float64   `json:"xCoord"`
	Y_Coord     float64   `json:"yCoord"`
}

// lightning strikes slice to seed lightning strike data.
var lightning_strikes = []lightning_strike{
	{ID: "1", Strike_Time: time.Date(2021, time.November, 10, 23, 0, 0, 0, time.UTC), X_Coord: 51.615335, Y_Coord: -1.193025},
	{ID: "2", Strike_Time: time.Date(2021, time.November, 10, 21, 0, 0, 0, time.UTC), X_Coord: 56.099230, Y_Coord: -3.922570},
	{ID: "3", Strike_Time: time.Date(2021, time.November, 10, 16, 0, 0, 0, time.UTC), X_Coord: 52.060492, Y_Coord: -4.638846},
}

func main() {
	router := gin.Default()
	router.GET("/strikes", getStrikes)
	router.POST("/strikes", postStrikes)

	router.Run("localhost:8080")
}

// getStrikes responds with the list of all lightning strikes as JSON.
func getStrikes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, lightning_strikes)
}

// postStrikes adds a lightning strike from JSON received in the request body.
func postStrikes(c *gin.Context) {
	var newStrike lightning_strike

	// Call BindJSON to bind the received JSON to
	// newStrike.
	if err := c.BindJSON(&newStrike); err != nil {
		return
	}

	// Add the new lightning strike to the slice.
	lightning_strikes = append(lightning_strikes, newStrike)
	c.IndentedJSON(http.StatusCreated, newStrike)
}
