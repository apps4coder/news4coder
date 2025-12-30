package cmd

import (
	"fmt"
	"news4coder/internal/official"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "列出所有官方新闻源",
	Long:  `显示所有可用的官方新闻源及其别名。`,
	Example: `  news4coder sources`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取官方源注册表
		registry := official.GetRegistry()
		sources := registry.List()

		if len(sources) == 0 {
			fmt.Println("暂无可用的官方新闻源")
			return nil
		}

		// 显示标题
		bold := color.New(color.Bold).SprintFunc()
		fmt.Println(bold("━━━ 官方新闻源列表 ━━━"))
		fmt.Println()

		// 显示表头
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s     %s\n", green("别名"), green("名称"))
		fmt.Println("────────────────────────────────────────────────────────")

		// 显示源列表
		blue := color.New(color.FgBlue).SprintFunc()
		for _, source := range sources {
			fmt.Printf("%-8s %s\n", blue(source.Alias), source.Name)
			if source.Description != "" {
				gray := color.New(color.FgHiBlack).SprintFunc()
				fmt.Printf("         %s\n", gray(source.Description))
			}
		}

		fmt.Println()
		fmt.Println(bold("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
		fmt.Println()

		// 使用提示
		gray := color.New(color.FgHiBlack).SprintFunc()
		fmt.Println(gray("💡 使用方法: news4coder <别名>"))
		fmt.Println(gray("💡 示例: news4coder infoq"))
		fmt.Println()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sourcesCmd)
}
