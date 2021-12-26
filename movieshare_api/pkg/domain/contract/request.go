package contract

type (
	// MoviePostRequestBody is request body for creating a movie.
	MoviePostRequestBody struct {
		UserID         string `json:"userId"`
		UserName       string `json:"userName"`
		Title          string `json:"title"`
		Overview       string `json:"overview"`
		Genre          string `json:"genre"`
		YouTubeLinkURL string `json:"youtubeLinkUrl"`
	}
)
