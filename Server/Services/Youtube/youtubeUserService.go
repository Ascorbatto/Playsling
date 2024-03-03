package youtubeservices

import (
	models "Conversify/Server/Models"
	services "Conversify/Server/Services"
	"net/http"
)

type YoutubeUserService struct {
	services.UserConfig
}

func NewYoutubeUserService(userConfig *services.UserConfig) *YoutubeUserService {
	return &YoutubeUserService{UserConfig: *userConfig}
}

var CurrentUser models.User

func (yus *YoutubeUserService) GetCurrentUser(w http.ResponseWriter, r *http.Request) {}
