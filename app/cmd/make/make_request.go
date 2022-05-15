package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)

	createFileFromStub(filePath, "request", model)
}