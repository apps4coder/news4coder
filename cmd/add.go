package cmd

import (
	"fmt"
	"news4coder/internal/storage"
	"news4coder/internal/subscription"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	addName string
	addURL  string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加新的订阅",
	Long:  `添加一个新的网站订阅，指定订阅名称和URL。`,
	Example: `  news4coder add --name "InfoQ中文站" --url "https://www.infoq.cn"
  news4coder add -n "Hacker News" -u "https://news.ycombinator.com"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 创建存储实例
		store, err := storage.New()
		if err != nil {
			return fmt.Errorf("初始化存储失败: %w", err)
		}

		// 加载现有配置
		config, err := store.Load()
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 创建订阅管理器
		manager := subscription.NewManager(config)

		// 添加订阅
		if err := manager.Add(addName, addURL); err != nil {
			return err
		}

		// 保存配置
		if err := store.Save(config); err != nil {
			return fmt.Errorf("保存配置失败: %w", err)
		}

		// 输出成功消息
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s 成功添加订阅：%s\n", green("✓"), addName)
		fmt.Printf("  URL: %s\n", addURL)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "订阅名称（必填）")
	addCmd.Flags().StringVarP(&addURL, "url", "u", "", "网站URL（必填）")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("url")
}
