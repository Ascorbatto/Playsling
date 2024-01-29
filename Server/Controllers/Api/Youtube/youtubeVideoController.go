package youtube

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	"encoding/json"
	"fmt"
)

var Songs []models.Song

func getVideoInfo(video_id string) models.VideoInfo {
	endpoint := fmt.Sprintf(utils.YoutubeVideoInfo + video_id)
	body, err := utils.GETRequest(endpoint, Token.AccessToken, nil)
	utils.ErrorManager(utils.ReadResponseError, err)

	var video models.VideoInfo
	err = json.Unmarshal(body, &video)
	utils.ErrorManager(utils.UnmarshalJSONError, err)

	return video
}

func GetSongs() {
	Songs = []models.Song{}
	for _, item := range PlaylistItems.Items {
		video_info := getVideoInfo(item.ContentDetails.VideoId)

		video := models.Song{
			Name:     utils.CleanString(video_info.Items[0].Snippet.Title),
			Author:   []string{utils.CleanAuthorString(video_info.Items[0].Snippet.ChannelTitle)},
			Duration: utils.ISO8601ToSec(video_info.Items[0].ContentDetails.Duration),
		}

		Songs = append(Songs, video)
	}
}
