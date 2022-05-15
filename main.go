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
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}