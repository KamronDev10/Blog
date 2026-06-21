package service

import (
	"blog_app/src/main/app/models"
	"blog_app/src/main/app/repository"
	"blog_app/src/main/common/token"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceI interface {
	CreateUser(user *models.User) (string, error)
	GetByEmail(email string) (*models.User, error)
	LogIn(email string, password string) (string, error)
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

func (us *userService) GetByEmail(email string) (*models.User, error) {

	user, err := us.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) LogIn(email string, password string) (string, error) {
	user, err := us.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user topilmadi ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return "", errors.New("parol noto'g'ri")
	}
	tokenstring, err := token.GetToken(user.Id, user.Username, user.Email, "user")
	if err != nil {
		return "", err

	}
	return tokenstring, nil
}
