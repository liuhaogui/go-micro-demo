package config

import "time"

type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetime() time.Duration
}

// mysql defaultMysqlConfig
type defaultMysqlConfig struct {
	URL               string        `json:"url"`
	Enable            bool          `json:"enabled"`
	MaxIdleConnection int           `json:"maxIdleConnection"`
	MaxOpenConnection int           `json:"maxOpenConnection"`
	ConnMaxLifetime   time.Duration `json:"connMaxLifetime"`
}

// URL mysql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

// Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}

// 连接数断开时间
func (m defaultMysqlConfig) GetConnMaxLifetime() time.Duration {
	return m.ConnMaxLifetime
}