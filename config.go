package goo_oceanengine

import "time"

type Config struct {
	AppId   string        `yaml:"app_id"`
	Secret  string        `yaml:"secret"`
	Debug   bool          `yaml:"debug"`
	Timeout time.Duration `yaml:"timeout"`
	Retry   int           `yaml:"retry"`
}
