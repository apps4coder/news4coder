# News4Coder 使用示例

本文档展示了 news4coder 命令行工具的常见使用场景。

## 场景 1: 快速开始

### 1. 添加你的第一个订阅

```bash
# 添加 InfoQ 中文站
.\news4coder.exe add --name "InfoQ中文站" --url "https://www.infoq.cn"
```

输出：
```
✓ 成功添加订阅：InfoQ中文站
  URL: https://www.infoq.cn
```

### 2. 查看订阅列表

```bash
.\news4coder.exe list
```

输出：
```
订阅列表：

序号   名称                   URL                                      创建时间
─────────────────────────────────────────────────────────────────────────────────────
1    InfoQ中文站             https://www.infoq.cn                     2025-12-14 01:45

总计: 1 个订阅
```

### 3. 获取最新内容（演示模式）

```bash
.\news4coder.exe fetch -n "InfoQ中文站" --demo
```

输出：
```
⟳ 正在获取 InfoQ中文站 的最新内容...

━━━ InfoQ中文站 最新内容 ━━━

1. Go 1.23 版本新特性详解
   🔗 https://www.infoq.cn/article/go-1.23-features
   本文详细介绍了 Go 1.23 的新特性，包括泛型改进、性能优化等内容...

2. 微服务架构下的分布式事务实践
   🔗 https://www.infoq.cn/article/distributed-transaction
   探讨在微服务架构中如何处理分布式事务的一致性问题...

[显示 10 条结果]
```

## 场景 2: 管理多个订阅源

### 添加多个技术网站

```bash
# 添加各种技术网站
.\news4coder.exe add -n "Hacker News" -u "https://news.ycombinator.com"
.\news4coder.exe add -n "GitHub Blog" -u "https://github.blog"
.\news4coder.exe add -n "Rust Blog" -u "https://blog.rust-lang.org"
.\news4coder.exe add -n "Go Blog" -u "https://go.dev/blog"
```

### 查看所有订阅

```bash
.\news4coder.exe list
```

## 场景 3: 删除订阅

### 按名称删除

```bash
.\news4coder.exe remove --name "Hacker News"
```

### 按序号删除

```bash
# 先查看列表获取序号
.\news4coder.exe list

# 删除第 3 个订阅
.\news4coder.exe remove --index 3
```

## 场景 4: 日常使用工作流

### 早晨查看技术新闻

```bash
# 查看 InfoQ 最新内容
.\news4coder.exe fetch -n "InfoQ中文站" -d

# 查看 Go 官方博客
.\news4coder.exe fetch -n "Go Blog" -d

# 查看 GitHub 博客
.\news4coder.exe fetch -n "GitHub Blog" -d
```

## 常见问题

### Q: 为什么推荐使用 --demo 参数？

A: 由于 Bing 搜索可能有反爬虫保护，实际搜索功能可能不稳定。演示模式可以让你查看工具的输出格式和功能演示。

### Q: 订阅数据保存在哪里？

A: Windows 系统保存在 `C:\Users\<用户名>\.news4coder\subscriptions.json`

### Q: 如何查看某个命令的帮助？

A: 使用 `--help` 参数，例如：

```bash
.\news4coder.exe fetch --help
```

### Q: 可以订阅哪些网站？

A: 任何 HTTP/HTTPS 的技术网站都可以，例如：
- InfoQ: https://www.infoq.cn
- Hacker News: https://news.ycombinator.com
- GitHub Blog: https://github.blog
- Rust Blog: https://blog.rust-lang.org
- Go Blog: https://go.dev/blog
- Python Blog: https://blog.python.org

## 提示和技巧

### 1. 使用短参数提高效率

```bash
# 长参数
.\news4coder.exe add --name "技术博客" --url "https://example.com"

# 短参数（更快）
.\news4coder.exe add -n "技术博客" -u "https://example.com"
```

### 2. 组合使用命令

```bash
# 添加后立即查看
.\news4coder.exe add -n "新网站" -u "https://example.com"; .\news4coder.exe list
```

### 3. 定期清理不常用的订阅

```bash
# 查看列表，找出不需要的订阅
.\news4coder.exe list

# 删除不需要的订阅
.\news4coder.exe remove -i 5
```

## 配置文件示例

订阅数据以 JSON 格式保存：

```json
{
  "subscriptions": [
    {
      "name": "InfoQ中文站",
      "url": "https://www.infoq.cn",
      "created_at": "2025-12-14T01:45:00Z"
    },
    {
      "name": "Go Blog",
      "url": "https://go.dev/blog",
      "created_at": "2025-12-14T01:50:00Z"
    }
  ]
}
```

你可以手动编辑此文件，但建议使用命令行工具进行管理。
