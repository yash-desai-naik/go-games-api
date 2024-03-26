// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Game struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	CurrentPrice  float64 `json:"currentPrice"`
	SellerName    string  `json:"sellerName"`
	DeveloperName string  `json:"developerName"`
	PublisherName string  `json:"publisherName"`
	ThumbnailURL  string  `json:"thumbnailURL"`
}

var games []Game

func main() {
	loadGamesFromJSON("games.json")

	router := gin.Default()

	router.GET("/", listGames)
	router.GET("/game", getGameByID)

	fmt.Println("Server listening on localhost:8080...")
	log.Fatal(router.Run(":8080"))
}

func listGames(c *gin.Context) {
	c.JSON(http.StatusOK, games)
}

func getGameByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, game := range games {
		if game.ID == id {
			c.JSON(http.StatusOK, game)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
}

func loadGamesFromJSON(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}

	err = json.Unmarshal(data, &games)
	if err != nil {
		log.Fatalf("Failed to parse JSON data: %v", err)
	}
}
