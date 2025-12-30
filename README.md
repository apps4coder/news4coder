# News4Coder - 程序员新闻订阅 CLI 工具

一个优雅、简单、易于维护的 Go 语言命令行工具，专为程序员设计的新闻订阅 CLI。

## 功能特性

- ✅ **订阅管理**：添加、查看、删除订阅源
- 🏷️ **别名支持**：为订阅设置短别名，快捷操作
- 🔍 **智能搜索**：使用 DuckDuckGo 站内搜索获取最新内容
- 🎯 **专注模式**：内置高质量官方信息源推荐
- 🎨 **优雅输出**：彩色终端输出，清晰易读
- 💾 **本地存储**：订阅数据保存在用户主目录
- 🚀 **快速启动**：单二进制文件，无需额外依赖

## 🎯 专注模式 - 官方信息源

专注模式提供高质量的官方信息源，直接从源站获取热点内容，无需搜索引擎中转。

| 别名 | 名称 | URL | 说明 |
|------|------|-----|------|
| `infoq` | InfoQ 中文站热点清单 | https://www.infoq.cn/hotlist | 软件开发、架构、AI、前后端技术资讯 |

### 专注模式 vs 普通模式

| 特性 | 专注模式（官方源） | 普通模式（用户订阅） |
|------|------------------|------------------|
| 数据来源 | 直接抓取原站热点内容 | DuckDuckGo 站内搜索 |
| 内容质量 | 精选热门文章 | 搜索引擎索引结果 |
| 支持源 | 内置官方源 | 任意网站 |
| 调用方式 | `news4coder infoq` | `fetch -n <别名>` |

### 使用专注模式

```bash
# 获取 InfoQ 中文站热点内容（专注模式）
.\news4coder.exe infoq

# 演示模式
.\news4coder.exe infoq --demo
```

输出示例：
```
🎯 ⟳ 专注模式 - 正在获取 InfoQ 中文站热点清单 的热点内容...

━━━ 🎯 InfoQ 中文站热点清单 热点内容 ━━━

1. 2025年技术趋势：AI原生应用的崛起
   🔗 https://www.infoq.cn/article/ai-native-apps-2025
   ...

🎯 专注模式：直接获取官方源 https://www.infoq.cn/hotlist
```

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
# 添加 InfoQ 中文站（带别名）
.\news4coder.exe add --name "InfoQ中文站" --alias infoq --url "https://www.infoq.cn"

# 或使用短参数
.\news4coder.exe add -n "Hacker News" -a hn -u "https://news.ycombinator.com"
```

#### 2. 查看订阅列表

```bash
.\news4coder.exe list
```

输出示例：
```
订阅列表：

序号   别名           名称                 URL                                 创建时间
───────────────────────────────────────────────────────────────────────────────────────────────
1    infoq        InfoQ中文站           https://www.infoq.cn                2025-12-14 01:45
2    hn           Hacker News          https://news.ycombinator.com        2025-12-14 02:00

总计: 2 个订阅
```

#### 3. 获取最新内容

```bash
# 专注模式 - 官方信息源（推荐）
.\news4coder.exe infoq

# 普通模式 - 使用别名获取内容
.\news4coder.exe fetch -n hn

# 演示模式
.\news4coder.exe infoq --demo
```

#### 4. 删除订阅

```bash
# 使用别名删除
.\news4coder.exe remove -n infoq

# 或按序号删除
.\news4coder.exe remove --index 1
```

## 命令参考

### `add` - 添加订阅

添加一个新的网站订阅。

**参数：**
- `--name, -n`：订阅名称（必填）
- `--alias, -a`：订阅别名/代号，用于快捷访问（可选）
- `--url, -u`：网站 URL（必填，必须是 HTTP/HTTPS 协议）

**示例：**
```bash
# 添加带别名的订阅
.\news4coder.exe add -n "技术博客" -a tech -u "https://example.com"

# 之后可用别名快捷操作
.\news4coder.exe fetch -n tech
```

### `list` - 列出订阅

显示所有已添加的订阅列表，包括序号、别名、名称、URL 和创建时间。

**示例：**
```bash
.\news4coder.exe list
```

### `infoq` - 专注模式

🎯 直接获取 InfoQ 中文站热点内容，无需搜索引擎中转。

**参数：**
- `--demo, -d`：演示模式，使用模拟数据（可选）

**示例：**
```bash
# 获取 InfoQ 热点内容
.\news4coder.exe infoq

# 演示模式
.\news4coder.exe infoq --demo
```

### `fetch` - 获取内容（普通模式）

获取用户订阅源的最新内容，使用 DuckDuckGo 站内搜索。

**参数：**
- `--name, -n`：订阅名称或别名（必填）
- `--demo, -d`：演示模式，使用模拟数据（可选）

**示例：**
```bash
# 使用别名获取内容
.\news4coder.exe fetch -n hn

# 使用完整名称获取内容
.\news4coder.exe fetch --name "Hacker News"

# 演示模式
.\news4coder.exe fetch -n hn --demo
```

**输出示例：**
```
⟳ 普通模式 - 正在搜索 Hacker News 的最新内容...

━━━ Hacker News 最新内容 ━━━

1. Go 1.23 版本新特性详解
   🔗 https://news.ycombinator.com/item?id=12345
   ...

━━━ 共 10 条结果 ━━━

💡 普通模式：基于 DuckDuckGo 站内搜索
```

### `remove` - 删除订阅

根据名称、别名或序号删除一个订阅。

**参数（二选一）：**
- `--name, -n`：订阅名称或别名
- `--index, -i`：订阅序号（从 1 开始）

**示例：**
```bash
# 使用别名删除
.\news4coder.exe remove -n tech

# 按序号删除
.\news4coder.exe remove -i 1
```

## 项目结构

```
news4coder/
├── cmd/                    # CLI 命令定义
│   ├── root.go            # 根命令
│   ├── infoq.go           # 🎯 专注模式命令
│   ├── add.go             # 添加订阅命令
│   ├── list.go            # 列出订阅命令
│   ├── remove.go          # 删除订阅命令
│   ├── fetch.go           # 获取内容命令
│   └── console_windows.go # Windows 控制台 UTF-8 支持
├── internal/              # 内部模块
│   ├── subscription/      # 订阅管理模块
│   │   ├── model.go       # 数据模型（含别名字段）
│   │   └── manager.go     # 订阅管理器
│   ├── official/          # 官方信息源模块（专注模式）
│   │   ├── model.go       # 官方源数据模型
│   │   ├── registry.go    # 官方源注册表
│   │   ├── fetcher.go     # 抓取器接口与工厂
│   │   └── infoq_fetcher.go # InfoQ 专用抓取器
│   ├── search/           # 搜索引擎模块（普通模式）
│   │   ├── model.go       # 搜索结果模型
│   │   └── engine.go      # DuckDuckGo 搜索引擎
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
      "alias": "infoq",
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

本工具使用 DuckDuckGo 搜索引擎的 HTML 版本。请注意：

1. **搜索结果**：结果取决于 DuckDuckGo 对网站的收录情况
2. **网络要求**：需要稳定的网络连接访问 DuckDuckGo
3. **超时设置**：默认超时时间为 30 秒

如果遇到搜索失败，可能的原因：
- 网站尚未被 DuckDuckGo 收录
- 网络连接问题
- 请求超时

**建议**：
- 使用 `--demo` 参数查看演示效果
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
