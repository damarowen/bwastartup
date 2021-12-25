package user

import (
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserService interface {
	RegisterUser(user DtoRegisterUserInput) (User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepo IUserRepository) IUserService {
	return &UserService{userRepo}
}

func (s *UserService) RegisterUser(input DtoRegisterUserInput) (User, error) {
	userToCreate := User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&input))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	userToCreate.PasswordHash = hashAndSalt([]byte(input.Password))
	userToCreate.Role = "user"
	data, errorSave := s.userRepository.Save(userToCreate)
	if errorSave != nil {
		log.Fatalf("Failed map %v", err)
	}
	return data, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
