package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	ArticleID int64  `gorm:"type:bigint(20); not null; foreignKey:ArticleID;" json:"article_id"`
	UserID    int64  `gorm:"type:bigint(20); not null; foreignKey:UserID;" json:"user_id"`
	Body      string `gorm:"type:varchar(256); not null" json:"body"`
	User      User
	CreatedAt time.Time      `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
