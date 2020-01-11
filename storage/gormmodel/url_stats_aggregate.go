package gormmodel

//URL Statistics
type URLStatsAggregate struct {
	UrlId    string `gorm:"column:url_id"`
	SumCount int    `gorm:"column:sum(count)"`
}
