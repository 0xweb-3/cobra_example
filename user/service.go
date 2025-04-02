package user

import (
	"context"
	"fmt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Start(ctx context.Context) {
	fmt.Println("User service is running and ready to serve requests...")
}
