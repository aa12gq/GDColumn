package migrations

import (
    "database/sql"
    "GDColumn/app/models"
    "GDColumn/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type Column struct {
        models.BaseModel

        AvatarID     uint64 `gorm:"type:bigint"`
        Author       uint64 `gorm:"type:bigint;`
        Title        string `gorm:"type:varchar(255);not null;index"`
        Description  string `gorm:"type:varchar(255);not null;index"`

        models.CommonTimestampsField
    }

    type Avatar struct {
        ID  uint64 `gorm:"type:bigint;"`
        URL string `gorm:"type:varchar(255);not null;index"`
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Column{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Column{})
    }

    migrate.Add("2022_05_16_230316_add_column_table", up, down)
}