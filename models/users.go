package models

import "time"

type User struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"type:varchar(128); not null" json:"username"`
	Email     string    `gorm:"type:varchar(128); not null" json:"email"`
	Password  string    `gorm:"type:varchar(128); not null" json:"password"`
	Salt      string    `gorm:"type:varchar(128); not null" json:"salt"`
	Name      string    `gorm:"type:varchar(128);" json:"name"`
	CreatedAt time.Time `gorm:"type:datetime; not null; default:CURRENT_TIMESTAMP" json:"created_at"`
}
