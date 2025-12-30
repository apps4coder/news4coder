package cmd

import (
	"fmt"
	"news4coder/internal/official"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "news4coder",
	Short: "程序员新闻订阅 CLI 工具",
	Long: `news4coder 是一个为程序员设计的新闻订阅命令行工具。
它可以帮助你订阅技术网站，并通过 Bing 站内搜索快速获取最新内容。

官方新闻源快捷访问:
  infoq       InfoQ 中文站热点清单

使用 "news4coder sources" 查看所有官方新闻源`,
	// 关闭默认的未知命令错误，允许自定义处理
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute 执行根命令
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// 检查是否为未知命令错误
		if strings.Contains(err.Error(), "unknown command") {
			// 提取命令名
			if len(os.Args) > 1 {
				alias := os.Args[1]
				// 尝试作为官方源别名处理
				if handleErr := handleOfficialSource(alias); handleErr == nil {
					return
				} else {
					// 如果官方源处理也失败，输出具体错误
					fmt.Fprintln(os.Stderr, handleErr)
					os.Exit(1)
					return
				}
			}
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// handleOfficialSource 处理官方源别名命令
func handleOfficialSource(alias string) error {
	// 获取官方源注册表
	registry := official.GetRegistry()
	source, exists := registry.Get(alias)
	if !exists {
		return fmt.Errorf("未知命令: %s\n\n运行 'news4coder --help' 查看可用命令", alias)
	}

	// 显示提示信息
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s 正在获取 %s 的最新内容...\n", cyan("⟳"), source.Name)
	fmt.Println()

	// 创建抓取器工厂
	factory := official.NewFetcherFactory()
	fetcher, err := factory.Create(source)
	if err != nil {
		return fmt.Errorf("创建抓取器失败: %w", err)
	}

	// 执行抓取
	results, err := fetcher.Fetch()
	if err != nil {
		return fmt.Errorf("获取内容失败: %w", err)
	}

	// 显示结果
	displayResults(results, source.Name)

	return nil
}
