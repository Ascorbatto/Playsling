package spotifyservices

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	services "Conversify/Server/Services"
	"encoding/json"
	"fmt"
	"net/http"
)

type SpotifyPlaylistService struct {
	services.PlaylistConfig
}

func NewSpotifyPlaylistService(playlistConfig *services.PlaylistConfig) *SpotifyPlaylistService {
	return &SpotifyPlaylistService{PlaylistConfig: *playlistConfig}
}

func (sps *SpotifyPlaylistService) GetPlaylistInfo(w http.ResponseWriter, r *http.Request, playlistId string) interface{} {

	endpoint := fmt.Sprintf(utils.SpotifyPlaylistItemsInfo + playlistId)

	body, err := utils.Request(sps.Client, "GET", endpoint, nil)
	//body, err := utils.GETRequest(endpoint, sps.PlaylistConfig.Client, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var PlaylistInfo models.PlaylistInfoS
	err = json.Unmarshal(body, &PlaylistInfo)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	return PlaylistInfo
}

func (sps *SpotifyPlaylistService) GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request) interface{} {

	//body, err := utils.GETRequest(utils.SpotifyCurrentUserPlaylists, sps.PlaylistConfig.Client, nil)
	body, err := utils.Request(sps.Client, "GET", utils.SpotifyCurrentUserPlaylists, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var CurrentUserPlaylistsS models.CurrentUserPlaylistsS
	err = json.Unmarshal(body, &CurrentUserPlaylistsS)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	return CurrentUserPlaylistsS
}
