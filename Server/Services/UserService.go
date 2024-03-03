package services

import "net/http"

type UserConfig struct {
	Client *http.Client
}

type UserService interface {
	GetCurrentUser(w http.ResponseWriter, r *http.Request)
}

func NewUserConfig(client *http.Client) *UserConfig {
	return &UserConfig{Client: client}
}
