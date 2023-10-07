package cmd

import (
	"github.com/ichaly/gcms/base"
	"github.com/ichaly/gcms/root"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"path/filepath"
)

var configFile string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start Service.",

	Run: func(cmd *cobra.Command, args []string) {
		if configFile == "" {
			configFile = filepath.Join("./conf", "dev.yml")
		}
		fx.New(
			root.Modules,
			//auth.Modules,
			base.Modules,
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
