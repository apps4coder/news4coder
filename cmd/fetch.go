package cmd

import (
	"fmt"
	"news4coder/internal/search"
	"news4coder/internal/storage"
	"news4coder/internal/subscription"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	fetchName string
	demoMode  bool
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "获取订阅的最新内容",
	Long:  `使用 Bing 搜索获取指定订阅源的最新内容（最多10条）。`,
	Example: `  news4coder fetch --name "InfoQ中文站"
  news4coder fetch -n "Hacker News"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fetchName == "" {
			return fmt.Errorf("请指定订阅名称（--name）")
		}

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

		// 获取订阅信息
		sub, err := manager.Get(fetchName)
		if err != nil {
			return err
		}

		// 显示提示信息
		cyan := color.New(color.FgCyan).SprintFunc()
		fmt.Printf("%s 正在获取 %s 的最新内容...\n", cyan("⟳"), sub.Name)
		fmt.Println()

		// 创建搜索引擎
		engine := search.NewEngine()

		// 执行搜索
		var results []search.SearchResult
		
		if demoMode {
			// 演示模式：使用模拟数据
			results = generateDemoResults(sub.Name, sub.URL)
		} else {
			// 正常模式：实际搜索
			var searchErr error
			results, searchErr = engine.Search(sub.URL)
			if searchErr != nil {
				return fmt.Errorf("搜索失败: %w\n\n提示：如果想查看演示效果，可以使用 --demo 参数", searchErr)
			}
		}

		// 显示结果
		displayResults(results, sub.Name)

		return nil
	},
}

// displayResults 格式化显示搜索结果
func displayResults(results []search.SearchResult, sourceName string) {
	bold := color.New(color.Bold).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(bold(fmt.Sprintf("━━━ %s 最新内容 ━━━", sourceName)))
	fmt.Println()

	for _, result := range results {
		// 序号和标题
		fmt.Printf("%s %s\n", green(fmt.Sprintf("%d.", result.Index)), bold(result.Title))

		// URL
		fmt.Printf("   %s %s\n", blue("🔗"), result.URL)

		// 摘要
		if result.Snippet != "" {
			// 截断过长的摘要
			snippet := result.Snippet
			if len(snippet) > 200 {
				snippet = snippet[:200] + "..."
			}
			fmt.Printf("   %s\n", wrapText(snippet, 80, "   "))
		}

		fmt.Println()
	}

	fmt.Println(bold(fmt.Sprintf("━━━ 共 %d 条结果 ━━━", len(results))))
	fmt.Println()

	// 提示信息
	gray := color.New(color.FgHiBlack).SprintFunc()
	fmt.Println(gray("💡 提示：以上结果基于 Bing 搜索"))
}

// wrapText 文本换行处理
func wrapText(text string, maxWidth int, indent string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var lines []string
	currentLine := indent

	for _, word := range words {
		if len(currentLine)+len(word)+1 <= maxWidth {
			if currentLine == indent {
				currentLine += word
			} else {
				currentLine += " " + word
			}
		} else {
			if currentLine != indent {
				lines = append(lines, currentLine)
			}
			currentLine = indent + word
		}
	}

	if currentLine != indent {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

// generateDemoResults 生成演示数据
func generateDemoResults(sourceName, sourceURL string) []search.SearchResult {
	return []search.SearchResult{
		{
			Index:   1,
			Title:   "Go 1.23 版本新特性详解",
			URL:     sourceURL + "/article/go-1.23-features",
			Snippet: "本文详细介绍了 Go 1.23 的新特性，包括泛型改进、性能优化等内容。新版本带来了更好的开发体验和更高的运行效率。",
		},
		{
			Index:   2,
			Title:   "微服务架构下的分布式事务实践",
			URL:     sourceURL + "/article/distributed-transaction",
			Snippet: "探讨在微服务架构中如何处理分布式事务的一致性问题，分享了多种解决方案和最佳实践。",
		},
		{
			Index:   3,
			Title:   "Kubernetes 1.29 新功能一览",
			URL:     sourceURL + "/article/kubernetes-1.29",
			Snippet: "Kubernetes 最新版本 1.29 发布，带来了更强大的容器编排功能和更好的安全性。",
		},
		{
			Index:   4,
			Title:   "Rust 在系统编程中的应用",
			URL:     sourceURL + "/article/rust-system-programming",
			Snippet: "介绍 Rust 语言在系统级编程中的优势，包括内存安全、并发处理等方面。",
		},
		{
			Index:   5,
			Title:   "前端性能优化最佳实践",
			URL:     sourceURL + "/article/frontend-performance",
			Snippet: "分享前端性能优化的各种技巧，包括资源加载、渲染优化、代码分割等方法。",
		},
		{
			Index:   6,
			Title:   "深入理解 Docker 容器技术",
			URL:     sourceURL + "/article/docker-deep-dive",
			Snippet: "从底层原理到实际应用，全面解析 Docker 容器技术，帮助开发者更好地使用容器化技术。",
		},
		{
			Index:   7,
			Title:   "AI 大模型应用开发指南",
			URL:     sourceURL + "/article/ai-llm-development",
			Snippet: "介绍如何利用大语言模型开发实际应用，包括 API 调用、提示工程等内容。",
		},
		{
			Index:   8,
			Title:   "PostgreSQL 高级特性与优化",
			URL:     sourceURL + "/article/postgresql-advanced",
			Snippet: "深入探讨 PostgreSQL 数据库的高级特性，包括查询优化、索引设计等。",
		},
		{
			Index:   9,
			Title:   "GraphQL 与 RESTful API 的选择",
			URL:     sourceURL + "/article/graphql-vs-rest",
			Snippet: "对比 GraphQL 和 RESTful API 的优缺点，帮助开发者选择适合的 API 设计方案。",
		},
		{
			Index:   10,
			Title:   "代码质量保障与自动化测试",
			URL:     sourceURL + "/article/code-quality-testing",
			Snippet: "讲解如何通过自动化测试和代码审查来提高代码质量，建立可靠的软件交付流程。",
		},
	}
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVarP(&fetchName, "name", "n", "", "订阅名称（必填）")
	fetchCmd.Flags().BoolVarP(&demoMode, "demo", "d", false, "演示模式（使用模拟数据）")
	fetchCmd.MarkFlagRequired("name")
}
