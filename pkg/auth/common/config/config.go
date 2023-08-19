package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AuthService  string `mapstructure:"Auth_SRV"`
	Port         string `mapstructure:"PORT"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var envs = []string{"AUTH_SRV", "PORT"}

// LoadConfig loads the configuration from file and environment variables
func LoadConfig() (*Config, error) {
	config := &Config{}
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return nil, fmt.Errorf("error binding env variable %s: %w", env, err)
		}
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey, err = generateSecretKey(32)
		if err != nil {
			return nil, fmt.Errorf("error generating secret key: %w", err)
		}
		config.JwtSecretKey = secretKey
	} else {
		config.JwtSecretKey = secretKey
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return config, nil
}

func generateSecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	secretKey := base64.URLEncoding.EncodeToString(randomBytes)
	return secretKey, nil
}

func (c *Config) GetJWTSecretKey() string {
	return c.JwtSecretKey

}
