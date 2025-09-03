package services

import (
	"v2/internal/modules/user/requests/auth"
	UserResponse "v2/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}
