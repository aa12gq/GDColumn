package main

import (
	"GDColumn/app/cmd"
	"GDColumn/app/cmd/make"
	"GDColumn/bootstrap"
	btsConfig "GDColumn/config"
	"GDColumn/pkg/config"
	"GDColumn/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	btsConfig.Initialize()
}
// @title 仿知乎专栏后端API
// @version 1.0
// @description GDColumn是一个基于 Gin 框架搭建的防知乎专栏后端 API。
// @termsOfService http://swagger.io/terms/

// @contact.name AA12
// @contact.url https://qm.qq.com/cgi-bin/qm/qr?k=qjar90G1gAtMF6uR7-8WLk5OvpD7sjUj&noverify=0
// @contact.email 875151567@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
// @schemes http https

func main() {

	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {

			config.InitConfig(cmd.Env)

			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
			bootstrap.SetupCache()
			bootstrap.SetupSnowflake()
			bootstrap.SetupOss()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}