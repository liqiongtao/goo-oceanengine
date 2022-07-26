package goo_oceanengine

import "time"

type Config struct {
	AppId   string        `yaml:"app_id" json:"app_id"`
	Secret  string        `yaml:"secret" json:"secret"`
	Debug   bool          `yaml:"debug" json:"debug"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
	Retry   int           `yaml:"retry" json:"retry"`
}
