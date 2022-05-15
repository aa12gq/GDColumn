package make

import (
	"fmt"
	"GDColumn/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller，exmaple: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	createFileFromStub(filePath, "apicontroller", model)
}