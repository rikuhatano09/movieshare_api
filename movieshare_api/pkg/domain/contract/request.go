package contract

type (
	// MoviePostRequestBody is request body for creating a movie.
	MoviePostRequestBody struct {
		UserID         string `json:"userId"`
		Title          string `json:"title"`
		Overview       string `json:"overview"`
		Genre          string `json:"genre"`
		YouTubeLinkURL string `json:"youtubeLinkUrl"`
	}
)
