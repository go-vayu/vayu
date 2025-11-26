package models

import "time"

type City struct {
	ID        int64     `json:"id"`
	Name      string    `xorm:"not null" json:"name"`
	State     string    `xorm:"not null" json:"state"`
	Latitude  float64   `xorm:"not null" json:"latitude"`
	Longitude float64   `xorm:"not null" json:"longitude"`
	CreatedAt time.Time `xorm:"timestampz not null created" json:"created_at"`
	UpdatedAt time.Time `xorm:"timestampz not null updated" json:"updated_at"`
}
