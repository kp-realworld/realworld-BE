package models

import "time"

// user 테이블과 1:N 관계를 맺는 Article 테이블의 모델
type Article struct {
	ID          int64      `gorm:"primary_key" json:"id"`
	UserID      int64      `gorm:"type:bigint(20); not null" json:"user_id"`
	Title       string     `gorm:"type:varchar(128); not null" json:"title"`
	Description string     `gorm:"type:varchar(128); not null" json:"description"`
	Body        string     `gorm:"type:varchar(128); not null" json:"body"`
	CreatedAt   time.Time  `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"type:datetime; null; default:null" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime; null; default:null" json:"deleted_at"`
}
