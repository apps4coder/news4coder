package cmd

import (
	"fmt"
	"news4coder/internal/storage"
	"news4coder/internal/subscription"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有订阅",
	Long:  `显示所有已添加的订阅列表。`,
	Example: `  news4coder list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 创建存储实例
		store, err := storage.New()
		if err != nil {
			return fmt.Errorf("初始化存储失败: %w", err)
		}

		// 加载配置
		config, err := store.Load()
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 创建订阅管理器
		manager := subscription.NewManager(config)
		subs := manager.List()

		// 检查是否为空
		if len(subs) == 0 {
			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Printf("%s 暂无订阅\n", yellow("!"))
			fmt.Println("使用 'news4coder add --name <名称> --url <URL>' 添加订阅")
			return nil
		}

		// 显示订阅列表
		bold := color.New(color.Bold).SprintFunc()
		fmt.Println(bold("订阅列表："))
		fmt.Println()

		// 表头
		fmt.Printf("%-4s %-20s %-40s %-20s\n", "序号", "名称", "URL", "创建时间")
		fmt.Println("─────────────────────────────────────────────────────────────────────────────────────")

		// 表内容
		for i, sub := range subs {
			fmt.Printf("%-4d %-20s %-40s %-20s\n",
				i+1,
				truncateString(sub.Name, 20),
				truncateString(sub.URL, 40),
				sub.CreatedAt.Format("2006-01-02 15:04"))
		}

		fmt.Println()
		fmt.Printf("总计: %d 个订阅\n", len(subs))

		return nil
	},
}

// truncateString 截断字符串，超过长度时添加省略号
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func init() {
	rootCmd.AddCommand(listCmd)
}
