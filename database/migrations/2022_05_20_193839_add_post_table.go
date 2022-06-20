package migrations

import (
    "GDColumn/app/models/column"
    "GDColumn/app/models/user"
    "database/sql"
    "GDColumn/app/models"
    "GDColumn/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type Post struct {
        models.BaseModel

        Title        string `gorm:"type:varchar(255);`
        Content      string `gorm:"type:varchar(255);not null;index"`
        Excerpt      string `gorm:"type:varchar(255);not null;index"`
        UserID       string `gorm:"type:bigint(20);not null;index"`
        Image        string `gorm:"type:varchar(255);not null;index"`
        ColumnID     string `gorm:"type:bigint(255);not null;index"`

        User         user.User
        Column       column.Column

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