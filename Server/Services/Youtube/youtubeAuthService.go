package youtubeservices

import (
	services "Conversify/Server/Services"
)

type YoutubeAuthService struct {
	services.AuthConfig
}

func NewYoutubeAuthService(authConfig *services.AuthConfig) *YoutubeAuthService {
	return &YoutubeAuthService{AuthConfig: *authConfig}
}
