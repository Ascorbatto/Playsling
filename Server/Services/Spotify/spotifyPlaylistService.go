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

var CurrentUserPlaylists models.CurrentUserPlaylistsS

func NewSpotifyPlaylistService(playlistConfig *services.PlaylistConfig) *SpotifyPlaylistService {
	return &SpotifyPlaylistService{PlaylistConfig: *playlistConfig}
}

func (sps *SpotifyPlaylistService) GetPlaylistInfo(w http.ResponseWriter, r *http.Request, playlistId string) {
	endpoint := fmt.Sprintf(utils.SpotifyPlaylistItemsInfo + playlistId)

	body, err := utils.Request("GET", sps.Client, endpoint, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var PlaylistInfo models.PlaylistInfoS
	err = json.Unmarshal(body, &PlaylistInfo)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(PlaylistInfo)
	utils.ErrorManager(utils.EncodingJSONError, err)
}

func (sps *SpotifyPlaylistService) GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request) {

	body, err := utils.Request("GET", sps.Client, utils.SpotifyCurrentUserPlaylists, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	err = json.Unmarshal(body, &CurrentUserPlaylists)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(CurrentUserPlaylists)
	utils.ErrorManager(utils.EncodingJSONError, err)
}

func (sps *SpotifyPlaylistService) GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request, playlistId string) {
	playlist_id := r.URL.Query().Get("playlist")
	endpoint := fmt.Sprintf(utils.SpotifyPlaylistItemsInfo+"%s/tracks", playlist_id)

	body, err := utils.Request("GET", sps.Client, endpoint, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var Tracks models.PlaylistTracks
	err = json.Unmarshal(body, &Tracks)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Tracks)
	utils.ErrorManager(utils.EncodingJSONError, err)
}

func (sps *SpotifyPlaylistService) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	if len(CurrentUserPlaylists.Items) == 0 {
		sps.GetCurrentUserPlaylists(w, r)
	}
	if len(CurrentUser.ID) == 0 {
		return
	}

}
