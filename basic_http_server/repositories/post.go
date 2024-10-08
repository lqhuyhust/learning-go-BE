package repositories

import (
	"httpServer/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

// SAve a post
func (r *PostRepository) Save(post *models.Post) error {
	return r.DB.Save(post).Error
}

// FindByID finds a post by its ID
func (p *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := p.DB.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// Find all posts
func (r *PostRepository) ShowPosts(page int, limit int) ([]models.Post, error) {
	var posts []models.Post
	if err := r.DB.Limit(limit).Offset((page - 1) * limit).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// Delete a post
func (r *PostRepository) Delete(postID uint) error {
	return r.DB.Where("id = ?", postID).Delete(&models.Post{}).Error
}
