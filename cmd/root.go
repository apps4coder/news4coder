package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "news4coder",
	Short: "程序员新闻订阅 CLI 工具",
	Long: `news4coder 是一个为程序员设计的新闻订阅命令行工具。
它可以帮助你订阅技术网站，并通过 Bing 站内搜索快速获取最新内容。`,
}

// Execute 执行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
