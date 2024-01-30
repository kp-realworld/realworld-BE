package models

import (
	"time"

	"gorm.io/gorm"
)

// user 테이블과 1:N 관계를 맺는 Article 테이블
type Article struct {
	ID            int64  `gorm:"primary_key" json:"id"`
	UserID        int64  `gorm:"type:bigint(20); not null; foreignKey:UserID;" json:"user_id"`
	Title         string `gorm:"type:varchar(128); not null" json:"title"`
	Description   string `gorm:"type:varchar(128); not null" json:"description"`
	Body          string `gorm:"type:varchar(128); not null" json:"body"`
	FavoriteCount int    `gorm:"type:int(11); not null; default:0" json:"favorite_count"`
	User          User
	//TagList       []ArticleTag  `gorm:"many2many:article_tags;"`
	Tags      []ArticleTag
	Likes     []ArticleLike  `gorm:"foreignkey:ArticleID;"`
	CreatedAt time.Time      `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ArticleTag struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	ArticleID int64          `gorm:"type:bigint(20); not null; foreignKey:ArticleID;" json:"article_id"`
	Tag       string         `gorm:"type:varchar(128); not null" json:"tag"`
	CreatedAt time.Time      `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//
//type ArticleTag struct {
//	ID        int64      `gorm:"primary_key" json:"id"`
//	ArticleID int64      `json:"article_id"`
//	Tag       string     `gorm:"type:varchar(128); not null" json:"tag"`
//	CreatedAt time.Time  `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
//	UpdatedAt *time.Time `gorm:"type:datetime; null; default:null" json:"updated_at"`
//	DeletedAt *time.Time `gorm:"type:datetime; null; default:null" json:"deleted_at"`
//}

type ArticleLike struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	UserID    int64          `gorm:"type:bigint(20); not null; foreignKey:UserID;" json:"user_id"`
	ArticleID int64          `gorm:"type:bigint(20); not null; foreignKey:ArticleID;" json:"article_id"`
	CreatedAt time.Time      `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
