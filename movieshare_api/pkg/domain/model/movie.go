package model

import (
	"time"
)

type (
	Movie struct {
		ID             uint64    `json:"id" gorm:"column:id;type:bigserial;not null;primary_key"`
		CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at;type:timestamp with time zone;not null"`
		UserID         string    `json:"userId" gorm:"column:user_id;type:text;not null"`
		UserName       string    `json:"userName" gorm:"column:user_name;type:text;not null"`
		Title          string    `json:"title" gorm:"column:title;type:text;not null"`
		Overview       string    `json:"overview" gorm:"column:overview;type:text;not null"`
		Genre          string    `json:"genre" gorm:"column:genre;type:text;not null"`
		YouTubeTitleID string    `json:"youtubeTitleId" gorm:"column:youtube_title_id;type:text;not null"`
		GrinningScore  *uint32   `json:"grinningScore" gorm:"column:grinning_score;type:integer"`
	}
)

func (Movie) TableName() string {
	return "movie"
}
