package seeders

import (
	"fmt"
	"GDColumn/app/models/user"
	"GDColumn/database/factories"
	"GDColumn/pkg/console"
	"GDColumn/pkg/logger"
	"GDColumn/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factories.MakeUsers(10)

		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}