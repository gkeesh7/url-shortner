package gormmodel

import "time"

func (a *Redirection) TableName() string {
	return "redirection"
}

type Redirection struct {
	ID        int32     `gorm:"column:id"`
	UrlId     string    `gorm:"column:url_id"`
	URL       string    `gorm:"column:url"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Expiry    time.Time `gorm:"column:expiry"`
}
