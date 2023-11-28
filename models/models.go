package models

import (
	"time"
)

type Trip struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"Title"`
	Desc        string    `json:"Description"`
	Url         string    `json:"URL"`
	Price       string    `json:"Price"`
	Capacity    int       `json:"Capacity"`
	Type        string    `json:"Type"`
	HasCompFood bool      `json:"hasCompFood"`
	Additional  string    `json:"additional"`
	ContactNo   string    `json:"ContactNo"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

type Get struct {
	ID uint `json:"id"`
}

type PaginationData struct {
	NextPage    int `json:"NextPage"`
	LastPage    int `json:"LastPage"`
	CurrentPage int `json:"CurrentPage"`
	TotalPage   int `json:"TotalPage"`
}
