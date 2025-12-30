package official

import "sync"

var (
	registry *Registry
	once     sync.Once
)

// Registry 官方源注册表，管理所有官方新闻源
type Registry struct {
	sources map[string]*Source // key 为别名
}

// GetRegistry 获取官方源注册表单例
func GetRegistry() *Registry {
	once.Do(func() {
		registry = &Registry{
			sources: make(map[string]*Source),
		}
		// 注册官方源
		registry.registerDefaultSources()
	})
	return registry
}

// registerDefaultSources 注册默认的官方源
func (r *Registry) registerDefaultSources() {
	// InfoQ 中文站热点清单
	r.sources["infoq"] = &Source{
		Alias:       "infoq",
		Name:        "InfoQ 中文站热点清单",
		URL:         "https://www.infoq.cn/hotlist",
		FetcherType: "infoq",
		Description: "InfoQ 中文站的热点文章列表",
		Enabled:     true,
	}
}

// Get 根据别名获取官方源
func (r *Registry) Get(alias string) (*Source, bool) {
	source, exists := r.sources[alias]
	if !exists || !source.Enabled {
		return nil, false
	}
	return source, true
}

// List 获取所有启用的官方源列表
func (r *Registry) List() []*Source {
	var sources []*Source
	for _, source := range r.sources {
		if source.Enabled {
			sources = append(sources, source)
		}
	}
	return sources
}
