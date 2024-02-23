package main

import groupiron "ironmanager/groupIron"

// var apiKey string = "MTA4ODk5NDcxNzUwMjU1ODM0OA.G_u5ak.ApFKSCggejECI268Ru61ntMtvhGKOVN6bPl0w4"

func main() {
	// groupiron.CallApi()
	players := groupIron.GetAllPlayersCurrentStatus(client)
	if players == nil {
		panic("Failed to get group players")
	}
	groupiron.GetPlayerCurrentStats(players, "BB Bright")

}
