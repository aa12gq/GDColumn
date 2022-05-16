package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example:  make seeder user",
	Run:   runMakeSeeder,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeSeeder(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.TableName)

	createFileFromStub(filePath, "seeder", model)
}