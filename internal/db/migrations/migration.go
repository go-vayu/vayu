// Package migrations handles database migrations.
package migrations

import (
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

var migrations = []*xormigrate.Migration{}

func Migrate(x *xorm.Engine) error {
	migrations = append(migrations, createCitiesTable0001)

	m := xormigrate.New(x, migrations)

	return m.Migrate()
}
