package config

type Config struct {
	MysqlConf *MysqlConfig `yaml:"mysql"`
}

// Config代表mysql实例的配置信息
type MysqlConfig struct {
	Driver       string `yaml:"driver"`
	DSN          string `yaml:"dsn"`
	DialTimeout  int    `yaml:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxLifeConns int    `yaml:"max_life_conns"`
	DebugSQL     bool   `yaml:"debug_sql"`
}