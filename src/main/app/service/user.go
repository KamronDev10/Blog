package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
	"blog_app/src/main/common/token"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceI interface {
	CreateUser(user *models.User) (string, error)
	// GetByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepoI
}

func NewUserService(userRepo repository.UserRepoI) UserServiceI {
	return &userService{userRepo: userRepo}
}

func (us *userService) CreateUser(user *models.User) (string, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password_hash), bcrypt.DefaultCost)
	user.Password_hash = string(hashedPassword)

	// 2. Bazaga saqlash
	err := us.userRepo.CreateUser(user)
	if err != nil {
		return "", err
	}

	tokenString, err := token.GetToken(user.Id, user.Username, user.Email, "user")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
