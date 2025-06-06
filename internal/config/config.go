// internal/config/config.go
package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env       string    `yaml:"env"`
	HTTP      HTTP      `yaml:"http"`
	Database  DBs       `yaml:"database"`
	Redis     Redis     `yaml:"redis"`
	Log       Log       `yaml:"log"`
	Auth      Auth      `yaml:"auth"`
	RateLimit RateLimit `yaml:"rate_limit"`
	Consul    Consul    `yaml:"consul"`
	Crypto    Crypto    `yaml:"crypto"`
}

type HTTP struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type DBs struct {
	Master DB `yaml:"master"`
	Slave  DB `yaml:"slave1"`
}

type DB struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Name            string        `yaml:"name"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type Redis struct {
	Addr     string        `yaml:"addr"`
	Password string        `yaml:"password"`
	DB       int           `yaml:"db"`
	PoolSize int           `yaml:"pool_size"`
	Timeout  time.Duration `yaml:"timeout"`
}

type Log struct {
	Level            string   `yaml:"level"`
	Encoding         string   `yaml:"encoding"`
	OutputPaths      []string `yaml:"output_paths"`
	ErrorOutputPaths []string `yaml:"error_output_paths"`
	MaxSize          int      `yaml:"max_size"`
	MaxBackups       int      `yaml:"max_backups"`
	MaxAge           int      `yaml:"max_age"`
	Compress         bool     `yaml:"compress"`
}

type Auth struct {
	SecretKey   string        `yaml:"secret_key"`
	TokenExpiry time.Duration `yaml:"token_expiry"`
}

type RateLimit struct {
	RPS   int `yaml:"rps"`
	Burst int `yaml:"burst"`
}

type Consul struct {
	Addr          string        `yaml:"addr"`
	ServiceID     string        `yaml:"service_id"`
	ServiceName   string        `yaml:"service_name"`
	CheckInterval time.Duration `yaml:"check_interval"`
}

type Crypto struct {
	AESKey string `yaml:"aes_key"`
}

func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config")

	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}
