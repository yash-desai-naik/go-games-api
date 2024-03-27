// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Game struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	CurrentPrice  float64 `json:"currentPrice"`
	SellerName    string  `json:"sellerName"`
	DeveloperName string  `json:"developerName"`
	PublisherName string  `json:"publisherName"`
	ThumbnailURL  string  `json:"thumbnailURL"`
}

type GameData struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []Game `json:"elements"`
			} `json:"searchStore"`
		} `json:"Catalog"`
	} `json:"data"`
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
	var gameList []gin.H
	for _, game := range games {
		gameList = append(gameList, gin.H{
			"id":          game.ID,
			"title":       game.Title,
			"description": game.Description,
		})
	}
	c.JSON(http.StatusOK, gameList)
}

func getGameByID(c *gin.Context) {
	id := c.Query("id")

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

	var gameData GameData
	err = json.Unmarshal(data, &gameData)
	if err != nil {
		log.Fatalf("Failed to parse JSON data: %v", err)
	}

	games = gameData.Data.Catalog.SearchStore.Elements
}
