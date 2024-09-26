package services

import (
	"httpServer/models"
	"httpServer/repositories"
)

type ReactionService struct {
	ReactionRepository *repositories.ReactionRepository
}

func NewReactionService(reactionRepository *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{
		ReactionRepository: reactionRepository,
	}
}

// Make reaction for post
func (s *ReactionService) MakeReaction(reactionTypeID, userID, postID uint) error {
	// check user make reaction before
	existingReaction, err := s.ReactionRepository.FindByUserAndPost(userID, postID)
	if err == nil && existingReaction != nil {
		existingReaction.ReactionTypeID = reactionTypeID
		return s.ReactionRepository.Save(existingReaction)
	}

	// save new reaction
	reaction := &models.Reaction{
		ReactionTypeID: reactionTypeID,
		UserID:         userID,
		PostID:         postID,
	}
	return s.ReactionRepository.Save(reaction)
}

// Delete reaction for post
func (s *ReactionService) DeleteReaction(userID, postID uint) error {
	// check user make reaction before
	existingReaction, err := s.ReactionRepository.FindByUserAndPost(userID, postID)
	if err != nil {
		return err
	}

	return s.ReactionRepository.Delete(existingReaction)
}
