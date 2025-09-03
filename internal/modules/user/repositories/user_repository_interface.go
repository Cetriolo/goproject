package repositories

import (
	userModel "v2/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
}
