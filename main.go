package main

import (
	"log"
	"github.com/Brightwater/ironmanager/groupIron"
)

func main() {
	// groupiron.CallApi()
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Couldn't load config", err)
	}
	client := groupIron.NewApiClient(config.GROUP_IRON_BASE_URL+"/get-group-data?from_time=1980-12-23T03:57:02.960Z", config.GROUP_IRON_TOKEN)
	players := groupIron.GetAllPlayersCurrentStatus(client)
	if players == nil {
		panic("Failed to get group players")
	}
	groupIron.GetPlayerCurrentStats(players, "BB Bright")

}
