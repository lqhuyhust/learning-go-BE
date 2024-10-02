package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"httpServer/config"
	"httpServer/models"
	"httpServer/repositories"
	"time"
)

type PostService struct {
	PostRepository *repositories.PostRepository
}

func NewPostService(postRepository *repositories.PostRepository) *PostService {
	return &PostService{
		PostRepository: postRepository,
	}
}

// Create a new post
func (s *PostService) CreatePost(userID uint, title string, content string) (*models.Post, error) {
	post := &models.Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	if err := s.PostRepository.Save(post); err != nil {
		return nil, err
	}
	return post, nil
}

// Show all posts of an author
func (s *PostService) ShowUserPosts(page int, limit int) ([]models.Post, error) {
	// make redis cacheKey
	cacheKey := fmt.Sprintf("user_posts_page_%d", page)

	// checck cacheKey exists
	cachedPosts, err := config.RedisPostClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var posts []models.Post
		json.Unmarshal([]byte(cachedPosts), &posts)
		return posts, nil
	}

	// if cacheKey not exists, query database
	var posts []models.Post
	posts, err = s.PostRepository.ShowPosts(page, limit)
	if err != nil {
		return nil, err
	}

	// store data to redis
	cacheData, _ := json.Marshal(posts)
	config.RedisPostClient.Set(context.Background(), cacheKey, cacheData, time.Minute*5)

	return posts, nil
}

// Show a post by its ID
func (s *PostService) ShowPostByID(id uint) (*models.Post, error) {
	post, err := s.PostRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// Update a post (only author can update)
func (s *PostService) UpdatePost(userID uint, id uint, title, content *string) (*models.Post, error) {
	post, err := s.PostRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Just author can update
	if post.UserID != userID {
		return nil, errors.New("you are not the author")
	}

	if title != nil {
		post.Title = *title
	}

	if content != nil {
		post.Content = *content
	}

	if err := s.PostRepository.Save(post); err != nil {
		return nil, err
	}

	return post, nil
}

// Delete a post (only author can delete)
func (s *PostService) DeletePostByID(id uint, userID uint) error {
	post, err := s.PostRepository.FindByID(id)
	if err != nil {
		return err
	}

	// Just author can delete
	if post.UserID != userID {
		return errors.New("you are not the author")
	}

	if err := s.PostRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
