package models

type Users struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(100); not null" json:"username"`
}
