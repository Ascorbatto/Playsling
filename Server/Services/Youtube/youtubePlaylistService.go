package youtubeservices

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	services "Conversify/Server/Services"
	"encoding/json"
	"fmt"
	"net/http"
)

type YoutubePlaylistService struct {
	services.PlaylistConfig
}

func NewYoutubePlaylistService(playlistConfig *services.PlaylistConfig) *YoutubePlaylistService {
	return &YoutubePlaylistService{PlaylistConfig: *playlistConfig}
}

func (yps *YoutubePlaylistService) GetPlaylistInfo(w http.ResponseWriter, r *http.Request, playlistId string) {
	endpoint := fmt.Sprint(utils.YoutubePlaylist + playlistId)
	//body, err := utils.GETRequest(endpoint, ypc.PlaylistConfig.Token.AccessToken, nil)
	body, err := utils.Request("GET", yps.Client, endpoint, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var PlaylistInfo models.PlaylistInfoYT
	err = json.Unmarshal(body, &PlaylistInfo)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

}

func (yps *YoutubePlaylistService) GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request) {}

func (yps *YoutubePlaylistService) GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request, playlistId string) {
}
