package service

import (
	"clase2/domain"
	"clase2/repository"
)

type UsersService interface {
	Create(user *domain.UsersPOSTBody) (*domain.UsersGETResponse, error)
	GetByID(userID string) (*domain.UsersGETResponse, error)
	DeleteByID(userID string) error
}

type usersService struct {
	usersRepo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersService{
		usersRepo: repo,
	}
}

func (s *usersService) Create(userBody *domain.UsersPOSTBody) (*domain.UsersGETResponse, error){
	var userRepo domain.UserRepository

	userRepo.FromPOSTBody(userBody)

	response, err := s.usersRepo.Create(&userRepo)

	if err != nil {
		return nil, err
	}

	var user domain.UsersGETResponse
	user.FromUserRepository(response)

	return &user, nil
}

func (s *usersService) GetByID(userID string) (*domain.UsersGETResponse, error){
	userRepo, err := s.usersRepo.GetByID(userID)

	if err != nil {
		return nil, err
	}
	if userRepo == nil {
		return nil, nil
	}

	var user domain.UsersGETResponse

	user.FromUserRepository(userRepo)

	return &user, nil
}
func (s *usersService) DeleteByID(userID string) error {
	return s.usersRepo.DeleteByID(userID)
}