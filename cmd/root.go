package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 注册单词格式转换工具
	rootCmd.AddCommand(wordCmd)

	// 注册时间转换工具
	rootCmd.AddCommand(timeCmd)

	// 注册数据库格式化工具
	rootCmd.AddCommand(sqlCmd)
}

