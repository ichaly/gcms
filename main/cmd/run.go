package cmd

import (
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/data"
	"github.com/ichaly/gcms/form"
	"github.com/ichaly/gcms/serv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"path/filepath"
)

var configFile string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "version subcommand show git version info.",

	Run: func(cmd *cobra.Command, args []string) {
		if configFile == "" {
			configFile = filepath.Join("../conf", "dev.yml")
		}
		fx.New(
			data.Modules,
			boot.Modules,
			serv.Modules,
			form.Modules,
			fx.Supply(configFile),
		).Run()
	},
}

func init() {
	runCmd.PersistentFlags().StringVarP(
		&configFile, "config", "c", "", "start app with config file",
	)
	rootCmd.AddCommand(runCmd)
}
