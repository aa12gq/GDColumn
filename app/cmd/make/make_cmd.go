package make

import (
	"GDColumn/pkg/console"
	"github.com/spf13/cobra"
	"fmt"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, exmaple: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	createFileFromStub(filePath, "cmd", model)

	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
