package model

type (
	Movie struct {
		ID                  uint64  `json:"id" gorm:"column:id;type:bigserial;not null;primary_key"`
		UserID              string  `json:"user_id" gorm:"column:user_id;type:text;not null"`
		Title               string  `json:"title" gorm:"column:title;type:text;not null"`
		Overview            string  `json:"overview" gorm:"column:overview;type:text;not null"`
		Genre               string  `json:"genre" gorm:"column:genre;type:text;not null"`
		YouTubeLinkURL      string  `json:"youtubeLinkUrl" gorm:"column:youtube_link_url;type:text;not null"`
		YouTubeThumbnailURL string  `json:"youtubeThumbnailUrl" gorm:"column:youtube_thumbnail_url;type:text;not null"`
		GrinningScore       *uint32 `json:"grinningScore" gorm:"column:grinning_score;type:integer"`
	}
)
