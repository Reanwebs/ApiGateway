package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AuthService          string `mapstructure:"Auth_SRV"`
	ConferenceService    string `mapstructure:"Conference_SRV"`
	VideoService         string `mapstructure:"Video_Service"`
	MonitaizationService string `mapstructure:"Monitaization_Service"`
	Port                 string `mapstructure:"PORT"`
	JwtSecretKey         string `mapstructure:"JWT_SECRET_KEY"`
	SwagTitle            string `mapstructure:"SwagTitle"`
	SwagDescription      string `mapstructure:"SwagDescription"`
	SwagVersion          string `mapstructure:"SwagVersion"`
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
	config.JwtSecretKey = secretKey

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return config, nil
}

// func generateSecretKey(length int) (string, error) {
// 	randomBytes := make([]byte, length)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	secretKey := base64.URLEncoding.EncodeToString(randomBytes)
// 	fmt.Println("secretKey\t", secretKey)
// 	return secretKey, nil
// }

func (c *Config) GetJWTSecretKey() string {
	return c.JwtSecretKey
}
