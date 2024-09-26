package repositories

import (
	"httpServer/models"
	"time"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	DB *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		DB: db,
	}
}

func (r *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return r.DB.Create(refreshToken).Error
}

func (r *RefreshTokenRepository) FindByRefreshToken(refreshToken string) (*models.RefreshToken, error) {
	var refreshTokenModel models.RefreshToken
	err := r.DB.Where("refresh_token = ?", refreshToken).First(&refreshTokenModel).Error
	if err != nil {
		return nil, err
	}
	return &refreshTokenModel, nil
}

func (r *RefreshTokenRepository) DeleteExpiredRefreshTokens() error {
	return r.DB.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{}).Error
}

func (r *RefreshTokenRepository) DeleteByUserID(userID uint) error {
	return r.DB.Unscoped().Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}

func (r *RefreshTokenRepository) Save(refreshToken *models.RefreshToken) error {
	return r.DB.Save(refreshToken).Error
}
