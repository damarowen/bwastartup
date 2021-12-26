package user

import (
	"errors"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserService interface {
	RegisterUser(user DtoRegisterUserInput) (User, error)
	LoginUser(user DtoLoginUserInput)(User, error)
	IsDuplicateEmail(email string) (bool, error)
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

func (s *UserService) LoginUser(input DtoLoginUserInput) (User,error){
	user, err := s.userRepository.FindByEmail(input.Email)

	if err != nil{
		return user, err
	}

	if user.ID == 0{
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return user, errors.New("Password not match")
	}

	return user, nil
}

func (s *UserService) IsDuplicateEmail(email string) (bool, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, err
	}
	return false, err
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
