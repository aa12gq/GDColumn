package make

import (
	"GDColumn/app/app"
	"GDColumn/pkg/console"
	"github.com/spf13/cobra"
	"fmt"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1),
}

func runMakeMigration(cmd *cobra.Command, args []string) {

	// 日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}

