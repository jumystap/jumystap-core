package service

import (
	"fmt"

	"github.com/jumystap/jumystap-core/internal/model"
	"github.com/jumystap/jumystap-core/internal/repository"
	"github.com/jumystap/jumystap-core/internal/utils"
)

type AuthService struct {
    repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
    return &AuthService{repo: repo}
}

func (s *AuthService) Login(email string, password string) (*model.User, error) {
    user, err := s.repo.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }

    hashedPassword, err := utils.HashUserPassword(password)
    if err != nil {
        return nil, err
    }
    
    if(hashedPassword == user.Password) {
        return user, nil
    }
    
    return nil, fmt.Errorf("User with %s email not found", email)
}
