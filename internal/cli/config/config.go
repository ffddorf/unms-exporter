// Package config implements CLI configuration management.
package config

import (
	"errors"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const envPrefix = "UNMS_EXPORTER"

var (
	DefaultServerAddress = "[::]:9806"
	DefaultLogLevel      = logrus.InfoLevel
)

type Config struct {
	ServerAddr   string       `mapstructure:"listen" split_words:"true"`
	LogLevel     logrus.Level `mapstructure:"log_level" split_words:"true"`
	TokenPerHost tokenMap     `mapstructure:"token" envconfig:"token"`
	ExtraMetrics []string     `mapstructure:"extras" envconfig:"extras"`
}

func New(args []string) (*Config, error) {
	conf := &Config{
		ServerAddr: DefaultServerAddress,
		LogLevel:   DefaultLogLevel,
	}
	if err := envconfig.Process(envPrefix, conf); err != nil {
		return nil, fmt.Errorf("invalid environment variables: %w", err)
	}

	flags := pflag.NewFlagSet("unms_exporter", pflag.ContinueOnError)
	flags.StringP("listen", "l", conf.ServerAddr, "Address for the exporter to listen on")
	flags.StringP("config", "c", "", "Config file to use")
	flags.StringSlice("extras", nil, "Enable additional metrics")
	if err := flags.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	if err := v.BindPFlags(flags); err != nil {
		// this should not happen
		return nil, fmt.Errorf("invalid config setup: %w", err)
	}

	if path := v.GetString("config"); path != "" {
		v.SetConfigFile(path)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("invalid config file %q: %w", path, err)
		}
	}
	if err := v.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("invalid command line flags: %w", err)
	}
	if err := conf.validate(); err != nil {
		return nil, fmt.Errorf("invalid config settings: %w", err)
	}

	return conf, nil
}

func (c *Config) validate() error {
	if len(c.TokenPerHost) < 1 {
		return errors.New("No token configured")
	}
	if c.ServerAddr == "" {
		return errors.New("Server addr can't be blank")
	}
	return nil
}
