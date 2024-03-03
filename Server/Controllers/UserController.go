package controllers

import (
	services "Conversify/Server/Services"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.GetCurrentUser(w, r)
}
