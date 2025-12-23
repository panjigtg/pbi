package config

import (
	// "fmt"
	// "os"
	"strings"
	"time"

	"github.com/spf13/viper"
)


type Config struct {
	App AppConfig `mapstructure:",squash"`
	DB  DBConfig  `mapstructure:",squash"`
	JWT JWTConfig `mapstructure:",squash"`
}

type AppConfig struct {
	Name string `mapstructure:"APP_NAME"` 
    Port int    `mapstructure:"APP_PORT"`
}

type DBConfig struct {
    User            string        `mapstructure:"DB_USER"`
    Password        string        `mapstructure:"DB_PASSWORD"`
    Host            string        `mapstructure:"DB_HOST"`
    Port            int           `mapstructure:"DB_PORT"`
    Name            string        `mapstructure:"DB_NAME"`
    MaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONN"`
    MaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONN"`
	ConnMaxLifetime int           `mapstructure:"DB_CONN_MAX_LIFETIME"` 
	ConnMaxIdleTime int           `mapstructure:"DB_CONN_MAX_IDLE_TIME"`
    ConnMaxLifetimeDur time.Duration 
    ConnMaxIdleTimeDur time.Duration 
}

type JWTConfig struct {
	Secret        string        `mapstructure:"JWT_SECRET"`
	ExpireMinutes int `mapstructure:"JWT_EXPIRE_MINUTES"`
	ExpireDur     time.Duration
}

func Load() (*Config, error) {
	// cwd, _ := os.Getwd()
	// fmt.Println("WORKDIR:", cwd)
	v := viper.New()
	
	v.SetConfigFile(".env")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// fmt.Printf("DB CFG: %+v\n", cfg.DB)

	cfg.DB.ConnMaxLifetimeDur = time.Duration(cfg.DB.ConnMaxLifetime) * time.Second
	cfg.DB.ConnMaxIdleTimeDur = time.Duration(cfg.DB.ConnMaxIdleTime) * time.Second
	cfg.JWT.ExpireDur = time.Duration(cfg.JWT.ExpireMinutes) * time.Minute

	return &cfg, nil
}
