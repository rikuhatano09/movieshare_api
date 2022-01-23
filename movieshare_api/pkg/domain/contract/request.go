package contract

type (
	// MoviePostRequestBody is request body for creating a movie.
	MoviePostRequestBody struct {
		UserID         string `json:"userId"`
		UserName       string `json:"userName"`
		Title          string `json:"title"`
		Overview       string `json:"overview"`
		Genre          string `json:"genre"`
		YouTubeTitleID string `json:"youtubeTitleId"`
	}

	// MoviePutRequestBody is request body for updating a movie.
	MoviePutRequestBody struct {
		GrinningScore *uint32 `json:"grinningScore"`
	}
)
