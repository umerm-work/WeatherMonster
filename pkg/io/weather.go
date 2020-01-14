package io

import (
	"time"
)

// Models is used to store time for sorting purpose
type Models struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

// City stores the location details of the city
type City struct {
	ID        int     `gorm:"column:id;primary_key" json:"id"`
	Name      string  `gorm:"column:name" json:"name"`
	Latitude  float64 `gorm:"column:latitude" json:"latitude"`
	Longitude float64 `gorm:"column:longitude" json:"longitude"`
	Models
}

// Temperatures stores temperature range of the city
type Temperatures struct {
	ID        int     `gorm:"column:id;primary_key" json:"id"`
	CityID    int     `gorm:"column:city_id;not null" json:"city_id"`
	Max       float64 `gorm:"column:max;not null" json:"max"`
	Min       float64 `gorm:"column:min;not null" json:"min"`
	Timestamp int64   `gorm:"column:timestamp;not null" json:"timestamp"`
	Models
}

// Forecast stores average value of min and max of city
type Forecast struct {
	ID     int     `gorm:"column:id;primary_key" json:"id,omitempty"`
	CityID int     `json:"city_id"`
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
	Sample int     `json:"sample"`
}
type WebHook struct {
	ID          int    `gorm:"column:id;primary_key" json:"id"`
	CityID      int    `gorm:"column:city_id;not null" json:"city_id"`
	CallbackUrl string `gorm:"column:callback_url;not null" json:"callback_url"`
}

// Response struct is for API response
type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success" `
	Error   string      `json:"error"`
}

// SuccessMessage helping function to make the success API response
func SuccessMessage(data interface{}) Response {
	return Response{
		Data:    data,
		Success: true,
	}
}

// FailureMessage helping function to make the failed API response
func FailureMessage(err string) Response {
	return Response{
		Success: false,
		Error:   err,
	}
}
