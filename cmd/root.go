package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	randomseed "goframe/submodule"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "简单的分布式服务器管理工具，支持minecraft-rcon.",
	Long:  "欲获取帮助， 请使用 --help 或 help。",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test small", randomseed.Small())
		fmt.Println("Test big", randomseed.Big())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
