package cmd

import (
	"errors"
	"fmt"
	"oms/cmd/app"
	"oms/cmd/jobs"
	"oms/config"
	"oms/pkg/global/log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "oms",
	Short:        "oms",
	SilenceUsage: true,
	Long:         `oms`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func tip() {
	usageStr := `使用 -h 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level, config.GetConfig().Application.Name)
	log.Logger.Info("init")
	rootCmd.AddCommand(jobs.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
