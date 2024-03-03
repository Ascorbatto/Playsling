package spotifyservices

import (
	services "Conversify/Server/Services"
)

type SpotifyAuthService struct {
	services.AuthConfig
}

func NewSpotifyAuthService(authConfig *services.AuthConfig) *SpotifyAuthService {
	return &SpotifyAuthService{AuthConfig: *authConfig}
}
