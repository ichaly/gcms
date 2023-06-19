package core

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
