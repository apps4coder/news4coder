package cmd

import (
	"fmt"
	"news4coder/internal/official"
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
	Long: `获取指定订阅源的最新内容。

专注模式：官方信息源（如 infoq）使用专用抓取器，直接获取原站热点内容。
普通模式：其他订阅源使用 DuckDuckGo 站内搜索获取内容。`,
	Example: `  # 专注模式 - 官方信息源
  news4coder fetch -n infoq
  
  # 普通模式 - 站内搜索
  news4coder fetch -n hn
  news4coder fetch --name "Hacker News"
  
  # 演示模式
  news4coder fetch -n infoq --demo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fetchName == "" {
			return fmt.Errorf("请指定订阅名称（--name）")
		}

		// 首先检查是否为官方信息源（专注模式）
		registry := official.GetRegistry()
		if source, exists := registry.Get(fetchName); exists {
			return fetchOfficialSource(source)
		}

		// 普通模式：从订阅列表中查找
		return fetchUserSubscription()
	},
}

// fetchOfficialSource 专注模式：获取官方信息源内容
func fetchOfficialSource(source *official.Source) error {
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()

	fmt.Printf("%s %s 专注模式 - 正在获取 %s 的热点内容...\n", magenta("🎯"), cyan("⟳"), source.Name)
	fmt.Println()

	if demoMode {
		// 演示模式
		results := generateDemoResults(source.Name, source.URL)
		displayOfficialResults(results, source.Name, source.URL)
		return nil
	}

	// 创建专用抓取器
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
	displayOfficialResults(results, source.Name, source.URL)
	return nil
}

// fetchUserSubscription 普通模式：获取用户订阅的内容
func fetchUserSubscription() error {
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
	fmt.Printf("%s 普通模式 - 正在搜索 %s 的最新内容...\n", cyan("⟳"), sub.Name)
	fmt.Println()

	// 执行搜索
	var results []search.SearchResult

	if demoMode {
		// 演示模式
		results = generateDemoResults(sub.Name, sub.URL)
	} else {
		// 创建搜索引擎
		engine := search.NewEngine()
		var searchErr error
		results, searchErr = engine.Search(sub.URL)
		if searchErr != nil {
			return fmt.Errorf("搜索失败: %w", searchErr)
		}
	}

	// 显示结果
	displayResults(results, sub.Name)
	return nil
}

// displayOfficialResults 显示官方信息源结果（专注模式）
func displayOfficialResults(results []search.SearchResult, sourceName, sourceURL string) {
	bold := color.New(color.Bold).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	fmt.Println(bold(fmt.Sprintf("━━━ 🎯 %s 热点内容 ━━━", sourceName)))
	fmt.Println()

	for _, result := range results {
		fmt.Printf("%s %s\n", green(fmt.Sprintf("%d.", result.Index)), bold(result.Title))
		fmt.Printf("   🔗 %s\n", makeClickableURL(result.URL))

		if result.Snippet != "" {
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

	fmt.Printf("%s 专注模式：直接获取官方源 %s\n", magenta("🎯"), makeClickableURL(sourceURL))
}

// displayResults 格式化显示搜索结果（普通模式）
func displayResults(results []search.SearchResult, sourceName string) {
	bold := color.New(color.Bold).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(bold(fmt.Sprintf("━━━ %s 最新内容 ━━━", sourceName)))
	fmt.Println()

	for _, result := range results {
		fmt.Printf("%s %s\n", green(fmt.Sprintf("%d.", result.Index)), bold(result.Title))
		fmt.Printf("   🔗 %s\n", makeClickableURL(result.URL))

		if result.Snippet != "" {
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

	gray := color.New(color.FgHiBlack).SprintFunc()
	fmt.Println(gray("💡 普通模式：基于 DuckDuckGo 站内搜索"))
}

// makeClickableURL 创建可点击的终端链接（使用 OSC 8 ANSI 转义序列）
func makeClickableURL(url string) string {
	// OSC 8 格式: \033]8;;URL\033\\TEXT\033]8;;\033\\
	// 这在支持 OSC 8 的终端中会创建可点击的超链接
	return fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, url)
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
	// 演示数据直接链接到源站
	return []search.SearchResult{
		{
			Index:   1,
			Title:   "Go 1.23 版本新特性详解",
			URL:     sourceURL,
			Snippet: "本文详细介绍了 Go 1.23 的新特性，包括泛型改进、性能优化等内容。新版本带来了更好的开发体验和更高的运行效率。",
		},
		{
			Index:   2,
			Title:   "微服务架构下的分布式事务实践",
			URL:     sourceURL,
			Snippet: "探讨在微服务架构中如何处理分布式事务的一致性问题，分享了多种解决方案和最佳实践。",
		},
		{
			Index:   3,
			Title:   "Kubernetes 1.29 新功能一览",
			URL:     sourceURL,
			Snippet: "Kubernetes 最新版本 1.29 发布，带来了更强大的容器编排功能和更好的安全性。",
		},
		{
			Index:   4,
			Title:   "Rust 在系统编程中的应用",
			URL:     sourceURL,
			Snippet: "介绍 Rust 语言在系统级编程中的优势，包括内存安全、并发处理等方面。",
		},
		{
			Index:   5,
			Title:   "前端性能优化最佳实践",
			URL:     sourceURL,
			Snippet: "分享前端性能优化的各种技巧，包括资源加载、渲染优化、代码分割等方法。",
		},
		{
			Index:   6,
			Title:   "深入理解 Docker 容器技术",
			URL:     sourceURL,
			Snippet: "从底层原理到实际应用，全面解析 Docker 容器技术，帮助开发者更好地使用容器化技术。",
		},
		{
			Index:   7,
			Title:   "AI 大模型应用开发指南",
			URL:     sourceURL,
			Snippet: "介绍如何利用大语言模型开发实际应用，包括 API 调用、提示工程等内容。",
		},
		{
			Index:   8,
			Title:   "PostgreSQL 高级特性与优化",
			URL:     sourceURL,
			Snippet: "深入探讨 PostgreSQL 数据库的高级特性，包括查询优化、索引设计等。",
		},
		{
			Index:   9,
			Title:   "GraphQL 与 RESTful API 的选择",
			URL:     sourceURL,
			Snippet: "对比 GraphQL 和 RESTful API 的优缺点，帮助开发者选择适合的 API 设计方案。",
		},
		{
			Index:   10,
			Title:   "代码质量保障与自动化测试",
			URL:     sourceURL,
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
