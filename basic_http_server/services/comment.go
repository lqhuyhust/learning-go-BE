package services

import (
	"errors"
	"httpServer/models"
	"httpServer/repositories"
)

type CommentService struct {
	CommentRepository *repositories.CommentRepository
}

func NewCommentService(commentRepository *repositories.CommentRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
	}
}

// Create a comment
func (s *CommentService) CreateComment(postID uint, userID uint, content string, parentID *uint) (*models.Comment, error) {
	comment := &models.Comment{
		Content:  content,
		UserID:   userID,
		PostID:   postID,
		ParentID: parentID,
	}

	if err := s.CommentRepository.Save(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// Get all comments for a post
func (s *CommentService) GetComments(postID uint) ([]models.Comment, error) {
	return s.CommentRepository.FindByPostID(postID)
}

// Update a comment (only author can update)
func (s *CommentService) UpdateComment(userID uint, commentID uint, content string) (*models.Comment, error) {
	comment, err := s.CommentRepository.FindByID(commentID)
	if err != nil {
		return nil, err
	}

	// Chỉ tác giả của comment mới có thể sửa
	if comment.UserID != userID {
		return nil, errors.New("you are not the author")
	}

	comment.Content = content

	if err := s.CommentRepository.Save(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// Delete a comment (only author can delete)
func (s *CommentService) DeleteComment(userID uint, commentID uint) error {
	comment, err := s.CommentRepository.FindByID(commentID)
	if err != nil {
		return err
	}

	// Chỉ tác giả của comment mới có thể xóa
	if comment.UserID != userID {
		return errors.New("you are not the author")
	}

	return s.CommentRepository.Delete(commentID)
}
