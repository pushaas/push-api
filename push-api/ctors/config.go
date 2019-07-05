package ctors

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	envVarName    = "PUSHAPI_ENV"
	configVarName = "PUSHAPI_CONFIG"
)

const defaultEnv = "local"

var envs = map[string]struct{}{
	defaultEnv: {},
	"prod": {},
}

func getEnvVariable() (string, error) {
	env := os.Getenv(envVarName)
	if env == "" {
		fmt.Println("[config] env variable not defined, falling back to default:", defaultEnv)
		env = defaultEnv
		return env, nil
	}

	if _, ok := envs[env]; !ok {
		return "", errors.New(fmt.Sprintf("you passed %s environment variable with an invalid value", envVarName))
	}

	fmt.Println("[config] env:", env)
	return env, nil
}

func setupFromConfigurationFile(config *viper.Viper, env string) error {
	// try to use custom config file, or falls back to file corresponding to env
	filepath := os.Getenv(configVarName)
	if filepath == "" {
		filepath = fmt.Sprintf("./config/%s.yml", env)
	}

	config.SetConfigFile(filepath)
	if err := config.ReadInConfig(); err != nil {
		if env == defaultEnv {
			fmt.Printf("[config] no config file found for default env in %s, using default config from code\n", filepath)
			return nil
		}
		return errors.New(fmt.Sprintf("error loading config file: %s", filepath))
	}

	fmt.Println("[config] loaded config from file:", filepath)
	return nil
}

func setupFromDefaults(config *viper.Viper, env string) {
	config.Set("env", env)

	// api
	config.SetDefault("api.enable_auth", true)
	config.SetDefault("api.statics_path", "./client/build")
	config.SetDefault("api.basic_auth_user", "app")

	// push_stream
	config.SetDefault("push_stream.host", "localhost")
	config.SetDefault("push_stream.port", "9080")

	// redis
	config.SetDefault("redis.db.channel.prefix", "ch")
	config.SetDefault("redis.db.stats_global.prefix", "stats_global")
	config.SetDefault("redis.db.stats_channel.prefix", "stats_channel")
	config.SetDefault("redis.pubsub.channels.publish", "publish")

	// server
	config.SetDefault("server.port", "8080")

	// workers
	config.SetDefault("workers.enabled", true)
	config.SetDefault("workers.persistent_channels.enabled", true)
	config.SetDefault("workers.persistent_channels.interval", "1m")
	config.SetDefault("workers.persistent_channels.lock_key", "lock_persistent_channels")
	config.SetDefault("workers.persistent_channels.lock_timeout", "50s")
}

func setupFromEnvironment(config *viper.Viper) {
	replacer := strings.NewReplacer(".", "__")
	config.SetEnvKeyReplacer(replacer)
	config.SetEnvPrefix("pushapi")
	config.AutomaticEnv()
}


func NewViper() (*viper.Viper, error) {
	env, err := getEnvVariable()
	if err != nil {
		return nil, err
	}

	config := viper.New()
	setupFromDefaults(config, env)
	if err := setupFromConfigurationFile(config, env); err != nil {
		return nil, err
	}
	setupFromEnvironment(config)

	return config, nil
}
