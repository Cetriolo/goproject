package repositories

import (
	userModel "v2/internal/modules/user/models"
	log "v2/logger"
	"v2/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (userRepository *UserRepository) Create(user userModel.User) userModel.User {

	var newUser userModel.User
	log.Log.Info("Creating user with email:", user.Name)
	result := userRepository.DB.Create(&user).Scan(&newUser)
	if result.Error != nil {
		log.Log.Error("Error creating user:", result.Error)
		return userModel.User{}
	}
	log.Log.Info("User created successfully with ID:", newUser.ID)
	return newUser
}

func (userRepository *UserRepository) FindByEmail(email string) userModel.User {
	var user userModel.User

	userRepository.DB.First(&user, "email = ?", email)

	return user

}
