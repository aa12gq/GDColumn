package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, exmaple: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeFactory(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	createFileFromStub(filePath, "factory", model)
}