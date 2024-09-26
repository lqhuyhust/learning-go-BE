package repositories

import (
	"httpServer/models"

	"gorm.io/gorm"
)

type ReactionRepository struct {
	DB *gorm.DB
}

func NewReactionRepository(db *gorm.DB) *ReactionRepository {
	return &ReactionRepository{
		DB: db,
	}
}

// Save a reaction
func (r *ReactionRepository) Save(reaction *models.Reaction) error {
	return r.DB.Save(reaction).Error
}

// Find a reaction
func (r *ReactionRepository) FindByUserAndPost(userID, postID uint) (*models.Reaction, error) {
	var reaction models.Reaction
	if err := r.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&reaction).Error; err != nil {
		return nil, err
	}
	return &reaction, nil
}

// Delete a reaction
func (r *ReactionRepository) Delete(reaction *models.Reaction) error {
	return r.DB.Delete(reaction).Error
}
