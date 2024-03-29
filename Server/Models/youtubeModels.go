package models

type PlaylistInfoYT struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
		Status struct {
			PrivacyStatus string `json:"privacyStatus"`
		} `json:"status"`
		ContentDetails struct {
			ItemCount uint `json:"itemCount"`
		} `json:"contentDetails"`
	}
}

type PlaylistItems struct {
	Items []struct {
		ContentDetails struct {
			VideoId string `json:"videoId"`
		} `json:"contentDetails"`
	} `json:"items"`
}

type VideoInfo struct {
	Items []struct {
		Snippet struct {
			Title        string `json:"title"`
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
		ContentDetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
	} `json:"items"`
}

type SearchVideos struct {
	Items []struct {
		Id struct {
			VideoId string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title        string `json:"title"`
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
	} `json:"items"`
}

type PostRequestBody struct {
	Snippet struct {
		PlaylistID string `json:"playlistId"`
		Position   int    `json:"position"`
		ResourceID struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"resourceId"`
	} `json:"snippet"`
}
