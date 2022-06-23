package migrations

import (
    "database/sql"
    "GDColumn/app/models"
    "GDColumn/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type Image struct {
        models.BaseModel

        URL          string `gorm:"type:varchar(255);not null;index"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Image{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Image{})
    }

    migrate.Add("2022_06_22_202819_add_image_table", up, down)
}