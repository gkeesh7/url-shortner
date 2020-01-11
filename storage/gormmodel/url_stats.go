package gormmodel

import "time"

func (a *URLStats) TableName() string {
	return "url_stats"
}

//URL Statistics
type URLStats struct {
	ID        int32     `gorm:"column:id"`
	UrlId     string    `gorm:"column:url_id"`
	Count     int32     `gorm:"column:count"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
