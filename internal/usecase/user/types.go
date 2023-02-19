package user

import (
	"go-otg/internal/entity"
)

type User entity.User

type LoginUserModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserModel struct {
	Username string `json:"username" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Age      int    `json:"age" binding:"required,max=3"`
}
