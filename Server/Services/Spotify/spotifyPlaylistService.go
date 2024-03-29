package spotifyservices

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	services "Conversify/Server/Services"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// TODO implement a way to verify if the current user ID and playlist are retrieved
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

func (sps *SpotifyPlaylistService) GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request, playlistID string) {
	endpoint := fmt.Sprintf(utils.SpotifyPlaylistItemsInfo+"%s/tracks", playlistID)

	body, err := utils.Request("GET", sps.Client, endpoint, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var Tracks models.PlaylistTracks
	err = json.Unmarshal(body, &Tracks)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Tracks)
	utils.ErrorManager(utils.EncodingJSONError, err)
}

func (sps *SpotifyPlaylistService) CreatePlaylist(w http.ResponseWriter, r *http.Request, playlistName string) {
	if len(CurrentUserPlaylists.Items) == 0 {
		sps.GetCurrentUserPlaylists(w, r)
	}
	if len(CurrentUser.ID) == 0 {
		w.Write([]byte("User ID not available."))
		return
	}
	for _, playlist := range CurrentUserPlaylists.Items {
		if strings.Contains(playlist.Name, playlistName) {
			w.Write([]byte("This playlist already exists."))
			log.Println(playlist.ID)
			sps.GetPlaylistItemsInfo(w, r, playlist.ID)
			return
		}
	}
	endpoint := fmt.Sprint(utils.SpotifyCreatePlaylist + CurrentUser.ID + "/playlists")
	log.Println(endpoint)

	data := map[string]string{
		"name": playlistName,
	}
	log.Println(data)

	var playlistInfo models.PlaylistInfoS

	jsonData, err := json.Marshal(data)
	log.Println(jsonData)
	utils.ErrorManager(utils.MarshalJSONError, err)

	body, err := utils.Request("POST", sps.Client, endpoint, bytes.NewBuffer(jsonData))
	log.Println(body)
	utils.ErrorManager(utils.ReadResponseError, err)

	err = json.Unmarshal(body, &playlistInfo)
	log.Println(playlistInfo)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(playlistInfo)
	utils.ErrorManager(utils.EncodingJSONError, err)
}
