package database

import (
	"httpServer/models"
)

func SeedReactions() {
	reactions := []string{
		"like",
		"love",
		"laugh",
		"wow",
		"sad",
		"angry",
	}

	for _, reaction := range reactions {
		reactionType := models.ReactionType{Name: reaction}
		DB.FirstOrCreate(&reactionType, models.ReactionType{Name: reaction})
	}
}
