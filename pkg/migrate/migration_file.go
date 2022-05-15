package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator, *sql.DB)

var migrationFiles []MigrationFile

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}