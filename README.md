# News4Coder - 程序员新闻订阅 CLI 工具

一个优雅、简单、易于维护的 Go 语言命令行工具，专为程序员设计的新闻订阅 CLI。

## 功能特性

- ✅ **订阅管理**：添加、查看、删除订阅源
- 🔍 **智能搜索**：使用 Bing 站内搜索获取最新内容
- 🎨 **优雅输出**：彩色终端输出，清晰易读
- 💾 **本地存储**：订阅数据保存在用户主目录
- 🚀 **快速启动**：单二进制文件，无需额外依赖

## 快速开始

### 安装

确保已安装 Go 1.18 或更高版本：

```bash
# 克隆仓库
git clone <repository-url>
cd news4coder

# 编译项目
go build -o news4coder.exe

# 或直接运行
go run main.go
```

### 基本使用

#### 1. 添加订阅

```bash
# 添加 InfoQ 中文站
.\news4coder.exe add --name "InfoQ中文站" --url "https://www.infoq.cn"

# 或使用短参数
.\news4coder.exe add -n "Hacker News" -u "https://news.ycombinator.com"
```

#### 2. 查看订阅列表

```bash
.\news4coder.exe list
```

输出示例：
```
订阅列表：

序号   名称                   URL                                      创建时间
─────────────────────────────────────────────────────────────────────────────────────
1    InfoQ中文站             https://www.infoq.cn                     2025-12-14 01:45

总计: 1 个订阅
```

#### 3. 获取最新内容

```bash
# 获取指定订阅的最新内容（最多10条）
.\news4coder.exe fetch --name "InfoQ中文站"

# 或使用短参数
.\news4coder.exe fetch -n "InfoQ中文站"

# 使用演示模式（显示模拟数据）
.\news4coder.exe fetch -n "InfoQ中文站" --demo
```

**注意**：由于 Bing 可能有反爬虫保护，实际搜索功能可能不稳定。建议使用 `--demo` 参数来查看演示效果。

#### 4. 删除订阅

```bash
# 按名称删除
.\news4coder.exe remove --name "InfoQ中文站"

# 按序号删除
.\news4coder.exe remove --index 1
```

## 命令参考

### `add` - 添加订阅

添加一个新的网站订阅。

**参数：**
- `--name, -n`：订阅名称（必填）
- `--url, -u`：网站 URL（必填，必须是 HTTP/HTTPS 协议）

**示例：**
```bash
.\news4coder.exe add -n "技术博客" -u "https://example.com"
```

### `list` - 列出订阅

显示所有已添加的订阅列表，包括序号、名称、URL 和创建时间。

**示例：**
```bash
.\news4coder.exe list
```

### `fetch` - 获取内容

使用 Bing 搜索获取指定订阅源的最新内容（最多 10 条结果）。

**参数：**
- `--name, -n`：订阅名称（必填）
- `--demo, -d`：演示模式，使用模拟数据（可选）

**示例：**
```bash
# 实际搜索
.\news4coder.exe fetch -n "InfoQ中文站"

# 演示模式
.\news4coder.exe fetch -n "InfoQ中文站" --demo
```

### `remove` - 删除订阅

根据名称或序号删除一个订阅。

**参数（二选一）：**
- `--name, -n`：订阅名称
- `--index, -i`：订阅序号（从 1 开始）

**示例：**
```bash
# 按名称删除
.\news4coder.exe remove -n "技术博客"

# 按序号删除
.\news4coder.exe remove -i 1
```

## 项目结构

```
news4coder/
├── cmd/                    # CLI 命令定义
│   ├── root.go            # 根命令
│   ├── add.go             # 添加订阅命令
│   ├── list.go            # 列出订阅命令
│   ├── remove.go          # 删除订阅命令
│   └── fetch.go           # 获取内容命令
├── internal/              # 内部模块
│   ├── subscription/      # 订阅管理模块
│   │   ├── model.go       # 数据模型
│   │   └── manager.go     # 订阅管理器
│   ├── search/           # 搜索引擎模块
│   │   ├── model.go       # 搜索结果模型
│   │   └── engine.go      # Bing 搜索引擎
│   └── storage/          # 存储模块
│       └── storage.go     # JSON 文件存储
├── main.go               # 程序入口
├── go.mod                # 依赖管理
└── README.md             # 项目说明
```

## 配置文件

订阅数据保存在：
- Windows: `C:\Users\<用户名>\.news4coder\subscriptions.json`
- macOS/Linux: `~/.news4coder/subscriptions.json`

配置文件示例：
```json
{
  "subscriptions": [
    {
      "name": "InfoQ中文站",
      "url": "https://www.infoq.cn",
      "created_at": "2025-12-14T01:45:00Z"
    }
  ]
}
```

## 技术栈

- **语言**：Go 1.25.5
- **CLI 框架**：[Cobra](https://github.com/spf13/cobra)
- **HTML 解析**：[goquery](https://github.com/PuerkitoBio/goquery)
- **彩色输出**：[color](https://github.com/fatih/color)

## 注意事项

### 关于搜索功能

本工具使用 Bing 搜索引擎的网页搜索功能。请注意：

1. **反爬虫保护**：Bing 可能会对频繁的自动化请求进行限制或要求验证
2. **搜索结果**：结果取决于 Bing 对网站的收录情况
3. **网络要求**：需要稳定的网络连接访问 Bing

如果遇到搜索失败，可能的原因：
- Bing 要求人机验证
- 网站尚未被 Bing 收录
- 网络连接问题

**建议**：
- 不要频繁调用 fetch 命令
- 使用时注意遵守 Bing 的使用条款
- 如需大量数据，建议使用官方 API

## 开发

### 构建

```bash
# 编译为可执行文件
go build -o news4coder.exe

# 跨平台编译
# Windows
GOOS=windows GOARCH=amd64 go build -o news4coder.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o news4coder

# Linux
GOOS=linux GOARCH=amd64 go build -o news4coder
```

### 运行测试

```bash
go test ./...
```

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

## 作者

news4coder 项目团队
