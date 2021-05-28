package types

type Config struct {
	Apps   []AppConfig `toml:"app"`
	Listen string
}

func (c Config) GetDefaultApp() AppConfig {
	var app AppConfig

	for _, a := range c.Apps {
		if a.Default {
			app = a
		}
	}

	return app
}

type AppConfig struct {
	Name    string
	Url     string
	Default bool
}
