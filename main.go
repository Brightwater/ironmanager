package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Brightwater/ironmanager/api"
	"github.com/Brightwater/ironmanager/groupIron"

)

func main() {

	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Couldn't load config", err)
	}

	client := groupIron.NewApiClient(config.GROUP_IRON_BASE_URL, config.GROUP_IRON_TOKEN)

	players := groupIron.GetAllPlayersCurrentStatus(client)
	if players == nil {
		panic("Failed to get group players")
	}
	
	router := api.SetupRoutes(client)

	fmt.Println("START API on port", config.PORT)
	log.Fatal(http.ListenAndServe(config.PORT, router))
}
