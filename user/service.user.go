package user

import (
	"errors"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserService interface {
	RegisterUser(user DtoRegisterUserInput) (User, error)
	LoginUser(user DtoLoginUserInput) (User, error)
	IsDuplicateEmail(email string) (bool, error)
	SaveAvatarUser(ID int, fileLocation string) (User, error)
	FindById(id int) (User, error)
	//GetAllUsers() ([]User, error)
	//UpdateUser(input FormUpdateUserInput) (User, error)
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

func (s *UserService) LoginUser(input DtoLoginUserInput) (User, error) {
	user, err := s.userRepository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
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

func (s *UserService) SaveAvatarUser(id int, fileLocation string) (User, error) {
	u, err := s.userRepository.FindById(id)
	if err != nil {
		return u, err
	}
	u.AvatarFileName = fileLocation
	u, err = s.userRepository.UpdateUser(u)
	if err != nil {
		return u, err
	}
	return u , nil

}

func (s *UserService) FindById (id int) (User, error) {
	u, err := s.userRepository.FindById(id)
	if err != nil {
		return u, err
	}
	return u , nil

}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

//func (s *service) GetAllUsers() ([]User, error) {
//	users, err := s.repository.FindAll()
//	if err != nil {
//		return users, err
//	}
//
//	return users, nil
//}
//
//func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
//	user, err := s.repository.FindByID(input.ID)
//	if err != nil {
//		return user, err
//	}
//
//	user.Name = input.Name
//	user.Email = input.Email
//	user.Occupation = input.Occupation
//
//	updatedUser, err := s.repository.Update(user)
//	if err != nil {
//		return updatedUser, err
//	}
//
//	return updatedUser, nil
//}