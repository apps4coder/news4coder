package search

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Engine 搜索引擎接口
type Engine struct {
	client *http.Client
}

// NewEngine 创建新的搜索引擎实例
func NewEngine() *Engine {
	return &Engine{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// extractDomain 从URL中提取域名
func extractDomain(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	
	// 返回主机名（例如：www.infoq.cn 或 infoq.cn）
	host := parsedURL.Host
	if host == "" {
		return "", fmt.Errorf("无法从URL提取域名")
	}
	
	return host, nil
}

// Search 使用Bing搜索指定网站的最新内容
func (e *Engine) Search(siteURL string) ([]SearchResult, error) {
	// 提取域名
	domain, err := extractDomain(siteURL)
	if err != nil {
		return nil, fmt.Errorf("URL解析失败: %w", err)
	}

	// 构造Bing搜索URL
	query := fmt.Sprintf("%s", domain)
	searchURL := fmt.Sprintf("https://www.bing.com/search?q=%s", url.QueryEscape(query))

	// 发送HTTP请求
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置User-Agent避免被识别为爬虫
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("网络请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("搜索请求失败，状态码: %d", resp.StatusCode)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTML解析失败: %w", err)
	}

	// 提取搜索结果
	results := []SearchResult{}
	index := 1

	// 尝试多种选择器以适配Bing的不同页面结构
	selectors := []string{
		"li.b_algo",
		".b_algo",
		"#b_results li",
		"ol#b_results li",
	}
	
	for _, selector := range selectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			if index > 10 {
				return
			}

			result := SearchResult{Index: index}

			// 提取标题和链接 - 尝试多种选择器
			titleElem := s.Find("h2 a")
			if titleElem.Length() == 0 {
				titleElem = s.Find("h2").Find("a")
			}
			if titleElem.Length() == 0 {
				titleElem = s.Find("a[href]")
			}
			
			result.Title = strings.TrimSpace(titleElem.Text())
			if href, exists := titleElem.Attr("href"); exists {
				// 过滤掉非目标网站的链接
				if strings.Contains(href, domain) || strings.HasPrefix(href, "http") {
					result.URL = href
				}
			}

			// 提取摘要 - 尝试多种选择器
			snippetElem := s.Find(".b_caption p")
			if snippetElem.Length() == 0 {
				snippetElem = s.Find("p")
			}
			if snippetElem.Length() == 0 {
				snippetElem = s.Find(".b_caption")
			}
			if snippetElem.Length() == 0 {
				snippetElem = s.Find("div.b_caption")
			}
			result.Snippet = strings.TrimSpace(snippetElem.First().Text())

			// 只添加有效的结果（至少有标题和URL）
			if result.Title != "" && result.URL != "" {
				results = append(results, result)
				index++
			}
		})
		
		// 如果已经找到结果，就不再尝试其他选择器
		if len(results) > 0 {
			break
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("未找到搜索结果，Bing可能需要验证或页面结构已变化。请尝试直接访问 https://www.bing.com/search?q=%s 查看搜索结果", url.QueryEscape(query))
	}

	return results, nil
}
