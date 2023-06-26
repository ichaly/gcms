package serv

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	App      *App        `mapstructure:"app" jsonschema:"title=App"`
	Cache    *DataSource `mapstructure:"cache" jsonschema:"title=Cache"`
	Database *DataSource `mapstructure:"database" jsonschema:"title=DataSource"`
}

type App struct {
	Name      string `mapstructure:"name" jsonschema:"title=Application Name"`
	Port      string `mapstructure:"port" jsonschema:"title=Application Port"`
	Host      string `mapstructure:"host" jsonschema:"title=Application Host"`
	Debug     bool   `mapstructure:"debug" jsonschema:"title=Debug"`
	Workspace string `mapstructure:"workspace" jsonschema:"title=root"`
}

type DataSource struct {
	Url      string       `mapstructure:"url"`
	Host     string       `mapstructure:"host"`
	Port     int          `mapstructure:"port"`
	Name     string       `mapstructure:"name"`
	Dialect  string       `mapstructure:"dialect"`
	Username string       `mapstructure:"username"`
	Password string       `mapstructure:"password"`
	Sources  []DataSource `mapstructure:"sources"`
	Replicas []DataSource `mapstructure:"replicas"`
}

func NewConfig(confPath string) (*Config, error) {
	return readInConfig(confPath)
}

func readInConfig(file string) (*Config, error) {
	cp := filepath.Dir(file)
	vi := newViper(cp, filepath.Base(file))

	if err := vi.ReadInConfig(); err != nil {
		return nil, err
	}

	if pcf := vi.GetString("inherits"); pcf != "" {
		cf := vi.ConfigFileUsed()
		vi = newViper(cp, pcf)

		if err := vi.ReadInConfig(); err != nil {
			return nil, err
		}

		if v := vi.GetString("inherits"); v != "" {
			return nil, fmt.Errorf("inherited config (%s) cannot itself inherit (%s)", pcf, v)
		}

		vi.SetConfigFile(cf)

		if err := vi.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "GJ_") || strings.HasPrefix(e, "SJ_") {
			kv := strings.SplitN(e, "=", 2)
			setKeyValue(vi, kv[0], kv[1])
		}
	}

	c := &Config{}

	if err := vi.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to decode config, %v", err)
	}

	return c, nil
}

func newViper(configPath, configFile string) *viper.Viper {
	vi := newViperWithDefaults()
	vi.SetConfigName(strings.TrimSuffix(configFile, filepath.Ext(configFile)))

	if configPath == "" {
		vi.AddConfigPath("./config")
	} else {
		vi.AddConfigPath(configPath)
	}

	return vi
}

func newViperWithDefaults() *viper.Viper {
	vi := viper.New()

	vi.SetDefault("debug", true)

	vi.SetDefault("app.port", "3000")

	vi.SetDefault("cache.dialect", "memory")

	vi.SetDefault("database.dialect", "postgres")
	vi.SetDefault("database.host", "localhost")
	vi.SetDefault("database.port", 5432)
	vi.SetDefault("database.username", "postgres")
	vi.SetDefault("database.password", "")
	vi.SetDefault("database.schema", "public")
	vi.SetDefault("database.pool_size", 10)

	vi.SetDefault("proxy.timeout", 1)

	vi.SetDefault("env", "development")

	_ = vi.BindEnv("env", "GO_ENV")
	_ = vi.BindEnv("host", "HOST")
	_ = vi.BindEnv("port", "PORT")

	return vi
}

func setKeyValue(vi *viper.Viper, key string, value interface{}) bool {
	if strings.HasPrefix(key, "GJ_") || strings.HasPrefix(key, "SG_") {
		key = key[3:]
	}
	uc := strings.Count(key, "_")
	k := strings.ToLower(key)

	if vi.Get(k) != nil {
		vi.Set(k, value)
		return true
	}

	for i := 0; i < uc; i++ {
		k = strings.Replace(k, "_", ".", 1)
		if vi.Get(k) != nil {
			vi.Set(k, value)
			return true
		}
	}

	return false
}
