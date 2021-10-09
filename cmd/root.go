package cmd

import "github.com/spf13/cobra"

// 根命令
var rootCmd = &cobra.Command{}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 注册单词命令
	rootCmd.AddCommand(wordCmd)
	// 注册时间命令
	rootCmd.AddCommand(timeCmd)
}
