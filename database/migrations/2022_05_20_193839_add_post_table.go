package migrations

import (
    "GDColumn/app/models"
    "GDColumn/pkg/migrate"
    "database/sql"

    "gorm.io/gorm"
)

func init() {

    type Post struct {
        models.BaseModel

        Title        string `gorm:"type:varchar(255);`
        Content      string `gorm:"type:varchar(21589);not null;index"`
        Excerpt      string `gorm:"type:varchar(21589);not null;index"`
        AuthorID     uint64 `gorm:"type:bigint(20);not null;index"`
        ImageID      uint64 `gorm:"type:bigint(20);not null;index"`
        ColumnID     uint64 `gorm:"type:bigint(20);not null;index"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Post{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Post{})
    }

    migrate.Add("2022_05_20_193839_add_post_table", up, down)
}