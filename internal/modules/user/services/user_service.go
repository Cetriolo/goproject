package services

//ArticleModel "v2/internal/modules/article/models"
import (
	"errors"
	UserModel "v2/internal/modules/user/models"
	UserRepository "v2/internal/modules/user/repositories"
	"v2/internal/modules/user/requests/auth"
	UserResponse "v2/internal/modules/user/responses"
	log "v2/logger"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (UserService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user UserModel.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		log.Log.Fatal("Hashing password failed:", err)
	}
	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)
	newUser := UserService.userRepository.Create(user)
	if newUser.ID == 0 {
		log.Log.Error("Failed to create user")
		return response, errors.New("Failed to create user")
	}

	return UserResponse.ToUser(newUser), nil
}

func (UserService *UserService) CheckUserExists(email string) bool {
	user := UserService.userRepository.FindByEmail(email)
	return user.ID != 0
}
