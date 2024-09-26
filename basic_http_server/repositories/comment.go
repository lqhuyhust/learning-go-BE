package repositories

import (
	"httpServer/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		DB: db,
	}
}

// Save a comment
func (r *CommentRepository) Save(comment *models.Comment) error {
	return r.DB.Save(comment).Error
}

// Find a comment by its ID
func (r *CommentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := r.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// Get all comments for a post
func (r *CommentRepository) FindByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// Delete a comment
func (r *CommentRepository) Delete(commentID uint) error {
	return r.DB.Where("id = ?", commentID).Delete(&models.Comment{}).Error
}
