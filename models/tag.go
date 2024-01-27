package models

import "time"

// user 테이블과 1:N 관계를 맺는 Article 테이블의 모델
type Tag struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	ArticleID int64      `json:"article_id"`
	tag       string     `gorm:"type:varchar(128); not null" json:"tag"`
	CreatedAt time.Time  `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime; null; default:null" json:"deleted_at"`
}
