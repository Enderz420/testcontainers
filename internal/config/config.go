package config

import (
	"fmt"
	"strings"

	"enderz.net/testcontainer-test/internal/database"
	"github.com/spf13/viper"
)

type Config struct {
	DB database.Config `json:"database"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	viper.SetDefault(
		"db.dsn",
		"sqlserver://sa:oPax9HFmjU4AVAqXEeA@database:1433?database=test&TrustServerCertificate=true",
	)
	viper.SetDefault("db.maxOpenConns", 256)
	viper.SetDefault("db.maxIdleConns", 256)
	viper.SetDefault("db.maxIdleTime", "15m")
	viper.SetDefault("db.timeout", 60)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("testcontainers")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg)
	
	return &cfg, nil
}
