package spotifyservices

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	services "Conversify/Server/Services"
	"encoding/json"
	"net/http"
)

type SpotifyUserService struct {
	services.UserConfig
}

var CurrentUser models.User

func NewSpotifyUserService(userConfig *services.UserConfig) *SpotifyUserService {
	return &SpotifyUserService{UserConfig: *userConfig}
}

func (sus *SpotifyUserService) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	body, err := utils.Request("GET", sus.Client, utils.SpotifyCurrentUser, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var CurrentUserS models.UserS
	err = json.Unmarshal(body, &CurrentUserS)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	CurrentUser = models.User{ID: CurrentUserS.ID, Name: CurrentUserS.DisplayName}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(CurrentUser)
	utils.ErrorManager(utils.EncodingJSONError, err)
}
