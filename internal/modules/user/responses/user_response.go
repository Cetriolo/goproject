package responses

import (
	"fmt"
	UserModels "v2/internal/modules/user/models"
)

type User struct {
	ID    uint
	Name  string
	Email string
	Image string
}

type Users struct {
	Data []User
}

func ToUser(user UserModels.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://i.pravatar.cc/150?u=%s", user.Name),
	}
}
