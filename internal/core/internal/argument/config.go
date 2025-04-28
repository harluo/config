package argument

import (
	"github.com/harluo/config/internal/core/internal/core"
)

type Config struct {
	path *core.Path
}

func newConfig(path *core.Path) *Config {
	return &Config{
		path: path,
	}
}

func (c *Config) Target() any {
	return c.path.Bind()
}

func (*Config) Name() string {
	return "config"
}

func (*Config) Aliases() []string {
	return []string{
		"c",
		"conf",
		"configuration",
	}
}

func (*Config) Usage() string {
	return "指定配置文件`路径`"
}
