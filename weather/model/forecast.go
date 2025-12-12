package model

import "time"

type ForecastData struct {
	Temperature float64 `gorm:"not null" json:"temperature"`
	WindSpeed   float64 `gorm:"not null" json:"wind_speed"`
	Humidity    int     `gorm:"not null" json:"humidity"`
	Clouds      int     `gorm:"not null" json:"clouds"`
	Latitude    float64 `gorm:"not null" json:"lat"`
	Longtitude  float64 `gorm:"not null" json:"long"`
}

type ForecastDB struct {
	ID           uint         `gorm:"primaryKey" json:"id"`
	ForecastData ForecastData `gorm:"embedded" `
	CreatedAt    time.Time    `gorm:"not null" json:"created_at"`
}

type WeatherNowReq struct {
	Latitude  float64 `query:"lat"  json:"lat"`
	Longitude float64 `query:"long" json:"long"`
}

type WeatherHistoryReq struct {
	DateFrom string `query:"from" json:"from"`
	DateTo   string `query:"to" json:"to"`
}
