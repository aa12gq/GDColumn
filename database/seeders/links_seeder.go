package seeders

import (
    "fmt"
    "GDColumn/database/factories"
    "GDColumn/pkg/console"
    "GDColumn/pkg/logger"
    "GDColumn/pkg/seed"

    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedLinksTable", func(db *gorm.DB) {

        links  := factories.MakeLinks(5)

        result := db.Table("links").Create(&links)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}