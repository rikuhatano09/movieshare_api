package model

type (
	Movie struct {
		Id             uint64 `json:"id" gorm:"column:id;type:bigserial;not null;primary_key"`
		Title          string `json:"title" gorm:"column:title;type:text;not null"`
		Overview       string `json:"overview" gorm:"column:overview;type:text;not null"`
		Genre          string `json:"genre" gorm:"column:genre;type:text;not null"`
		YouTubeLinkURL string `json:"youtubeLinkUrl" gorm:"column:youtube_link_url;type:text;not null"`
	}
)
