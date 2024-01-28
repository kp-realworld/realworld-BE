package models

import "time"

type ArticleTag struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	ArticleID int64      `json:"article_id"`
	Tag       string     `gorm:"type:varchar(128); not null" json:"tag"`
	CreatedAt time.Time  `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime; null; default:null" json:"deleted_at"`
}
