package setup

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"db"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"vinyl-analytics"`
	DBUser     string `envconfig:"DB_USER" default:"vinyl"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"secret"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`

	DiscogsToken string `envconfig:"DISCOGS_TOKEN" default:""`
}

func Load() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("VA", c)
	return c, err
}

func BuildConfig() (*Config, error) {

	config, err := Load()
	if err != nil {
		return nil, err
	}
	return config, nil
}
