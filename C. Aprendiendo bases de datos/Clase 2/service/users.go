package service

import (
	"clase2/domain"
	"clase2/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type UsersService interface {
	Create(user *domain.UsersPOSTBody) *domain.UsersGETResponse
	GetByID(userID string) *domain.UsersGETResponse
}

type usersService struct {
	usersRepo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersService{
		usersRepo: repo,
	}
}

func (s *usersService) Create(userBody*domain.UsersPOSTBody) *domain.UsersGETResponse{
	var userRepo domain.UserRepository

	userRepo.FromPOSTBody(userBody)

	// Genera un UUID
	uuid, err := uuid.New()
	if err != nil {
		panic(err)
	}

	userRepo.ID = fmt.Sprintf("%X", uuid)

	userRepo = *s.usersRepo.Create(&userRepo)

	var user domain.UsersGETResponse

	user.FromUserRepository(&userRepo)

	return &user
}

func (s *usersService) GetByID(userID string) *domain.UsersGETResponse {
	userRepo := s.usersRepo.GetByID(userID)

	var user domain.UsersGETResponse

	user.FromUserRepository(userRepo)

	return &user
}