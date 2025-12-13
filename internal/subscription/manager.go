package subscription

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Manager 提供订阅管理功能
type Manager struct {
	config *Config
}

// NewManager 创建订阅管理器
func NewManager(config *Config) *Manager {
	return &Manager{config: config}
}

// GetConfig 获取当前配置
func (m *Manager) GetConfig() *Config {
	return m.config
}

// Add 添加新订阅
func (m *Manager) Add(name, urlStr string) error {
	// 验证名称
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("订阅名称不能为空")
	}

	if len(name) > 50 {
		return fmt.Errorf("订阅名称长度不能超过50个字符")
	}

	// 验证URL格式
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("URL格式无效: %w", err)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("URL必须是HTTP或HTTPS协议")
	}

	// 检查名称是否已存在
	for _, sub := range m.config.Subscriptions {
		if sub.Name == name {
			return fmt.Errorf("订阅名称已存在: %s", name)
		}
	}

	// 添加新订阅
	newSub := Subscription{
		Name:      name,
		URL:       urlStr,
		CreatedAt: time.Now(),
	}

	m.config.Subscriptions = append(m.config.Subscriptions, newSub)
	return nil
}

// Remove 删除订阅（按名称）
func (m *Manager) Remove(name string) error {
	for i, sub := range m.config.Subscriptions {
		if sub.Name == name {
			// 删除该订阅
			m.config.Subscriptions = append(m.config.Subscriptions[:i], m.config.Subscriptions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("订阅不存在: %s", name)
}

// RemoveByIndex 删除订阅（按序号，从1开始）
func (m *Manager) RemoveByIndex(index int) error {
	if index < 1 || index > len(m.config.Subscriptions) {
		return fmt.Errorf("序号无效: %d", index)
	}

	m.config.Subscriptions = append(m.config.Subscriptions[:index-1], m.config.Subscriptions[index:]...)
	return nil
}

// List 获取所有订阅
func (m *Manager) List() []Subscription {
	return m.config.Subscriptions
}

// Get 根据名称获取订阅
func (m *Manager) Get(name string) (*Subscription, error) {
	for _, sub := range m.config.Subscriptions {
		if sub.Name == name {
			return &sub, nil
		}
	}
	return nil, fmt.Errorf("订阅不存在: %s", name)
}
