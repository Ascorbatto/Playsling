package controllers

import (
	services "Conversify/Server/Services"
	"net/http"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ac.AuthService.HandleLogin(w, r)
}

func (ac *AuthController) HandleCallback(w http.ResponseWriter, r *http.Request) *http.Client {
	client := ac.AuthService.HandleCallback(w, r)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	return client
}
