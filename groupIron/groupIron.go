package groupIron

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
)

var levelLookup map[int]int64 = make(map[int]int64)

var skillsMapping = [23]string{"Agility", "Attack", "Construction", "Cooking", "Crafting",
	"Defence", "Farming", "Firemaking", "Fishing", "Fletching", "Herblore", "Hitpoints", "Hunter", "Magic",
	"Mining", "Prayer", "Ranged", "Runecraft", "Slayer", "Smithing", "Strength", "Thieving", "Woodcutting"}

func init() {
	// Precalculate XP for levels and store in levelLookup map
	for i := 1; i <= 99; i++ {
		levelLookup[i+1] = xpForLevel(i)
	}
}

type PlayerUnconverted struct {
	Name   string  `json:"name"`
	Skills []int64 `json:"skills"`
}

type Skill struct {
	Name            string
	Level           int16
	Xp              int64
	ProgressPercent int16
}

type Player struct {
	Name   string
	Skills map[string]Skill
}

func GetAllPlayersCurrentStatus(client *ApiClient) *[]Player {

	jsonpayload, err := client.GetData()
	if err != nil {
		log.Println(err)
		return nil
	}

	var allPlayersUnconverted []PlayerUnconverted
	err = json.Unmarshal([]byte(jsonpayload), &allPlayersUnconverted)
	if err != nil {
		fmt.Printf("Error happened unmarshal %s", err)
		return nil
	}

	allPlayers := mapAllPlayers(&allPlayersUnconverted)

	return allPlayers
}

func mapAllPlayers(allPlayersUnconverted *[]PlayerUnconverted) *[]Player {
	allPlayers := []Player{}

	for _, unmappedPlayer := range *allPlayersUnconverted {
		if len(unmappedPlayer.Skills) <= 0 {
			continue
		}

		playerSkills := mapPlayer(&unmappedPlayer)

		allPlayers = append(allPlayers, Player{unmappedPlayer.Name, playerSkills})
	}
	return &allPlayers
}

func mapPlayer(unmappedPlayer *PlayerUnconverted) map[string]Skill {
	playerSkills := make(map[string]Skill)
	for i, unmappedSkill := range unmappedPlayer.Skills {

		skill := Skill{
			skillsMapping[i],
			int16(calculateLevel(unmappedSkill)),
			unmappedSkill,
			0,
		}
		applyProgressPercent(&skill)

		playerSkills[skillsMapping[i]] = skill
	}

	return playerSkills
}

func GetPlayerCurrentStats(players *[]Player, playerName string) (*Player, error) {

	fmt.Println("Player to find", playerName)
	for _, p := range *players {
		if p.Name == playerName {
			player := &p
			fmt.Println(*player)
			return player, nil
		}
	}

	return nil, errors.New("Couldn't find player")
}

func applyProgressPercent(skill *Skill) {
	nextLevel := skill.Level + 1
	thisLevelXp := xpForLevel(int(skill.Level - 1))
	xpNeededToHitNextLevel := xpForLevel(int(nextLevel-1)) - thisLevelXp

	progressXp := skill.Xp - thisLevelXp

	// prevent divide by 0 panic
	if xpNeededToHitNextLevel == 0 {
		return
	}
	skill.ProgressPercent = int16(math.Floor(float64(progressXp) / float64(xpNeededToHitNextLevel) * 100))
}

func xpForLevel(level int) int64 {
	var xp int64 = 0
	for i := 1; i <= level; i++ {
		xp += int64(math.Floor(float64(i) + 300.0*math.Pow(2.0, float64(i)/7.0)))
	}
	return int64(math.Floor(0.25 * float64(xp)))
}

func calculateLevel(xp int64) int {
	for i := 1; i <= 99; i++ {
		start := levelLookup[i]
		end := levelLookup[i+1]
		if xp >= start && xp < end {
			return i
		}
	}
	return 99
}
