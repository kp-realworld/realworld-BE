package models

import "time"

type User struct {
	UserID        int64      `gorm:"primary_key" json:"user_id"`
	Username      string     `gorm:"type:varchar(64); not null" json:"username"`
	Email         string     `gorm:"type:varchar(64); not null" json:"email"`
	Password      string     `gorm:"type:varchar(128); not null" json:"password"`
	Bio           *string    `gorm:"type:varchar(128); null; default:null" json:"name"`
	ProfileImage  string     `gorm:"type:varchar(128);" json:"profile_image"`
	Articles      []Article  `gorm:"foreignkey:UserID"`
	LikedArticles []Article  `gorm:"many2many:article_likes;"`
	CreatedAt     time.Time  `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"type:datetime; null; default:null" json:"deleted_at"`
}
