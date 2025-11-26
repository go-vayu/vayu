package migrations

import (
	"time"

	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

type city0001 struct {
	ID        int64     `json:"id"`
	Name      string    `xorm:"not null" json:"name"`
	State     string    `xorm:"not null" json:"state"`
	Latitude  float64   `xorm:"not null" json:"latitude"`
	Longitude float64   `xorm:"not null" json:"longitude"`
	CreatedAt time.Time `xorm:"timestampz not null created" json:"created_at"`
	UpdatedAt time.Time `xorm:"timestampz not null updated" json:"updated_at"`
}

func (t *city0001) TableName() string {
	return "cities"
}

var createCitiesTable0001 = &xormigrate.Migration{
	ID:          "0001",
	Description: "Create the cities table",
	Migrate: func(e *xorm.Engine) error {
		return e.Sync(&city0001{})
	},
	Rollback: func(e *xorm.Engine) error {
		return e.DropTables(&city0001{})
	},
}
