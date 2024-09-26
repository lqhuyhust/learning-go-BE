package repositories

import (
	"httpServer/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := u.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Save(user *models.User) error {
	return u.DB.Save(user).Error
}
