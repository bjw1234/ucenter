package config

const (
	Prod  = "prod"
	Stage = "stage"
	Dev   = "dev"
)

type redisConfig struct {
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	PoolSize    int    `yaml:"pool_size"`
	IdleTimeout int    `yaml:"idle_timeout"`
	Retries     int    `yaml:"retries"`
}

type mysqlConfig struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	DataBase        string `yaml:"database"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Timezone        string `yaml:"timezone"`
	Timeout         int    `yaml:"timeout"`
	ReadTimeout     int    `yaml:"read_time_out"`
	WriteTimeout    int    `yaml:"write_time_out"`

	ConnMaxLifeTime int    `yaml:"conn_max_life_time"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
}

type Model struct {
	Env     string      `yaml:"env"`
	Listen  string      `yaml:"listen"`
	HomeDir string      `yaml:"home_dir"`
	Redis   redisConfig `yaml:"redis"`
	MySql   mysqlConfig `yaml:"mysql"`
}
