package context

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type logging struct {
	Debug     bool   `mapstructure:"debug"`
	Formatter string `mapstructure:"formatter"`
	Level     string `mapstructure:"level"`
}

// Config defines the application configuration
type Config struct {
	Host      string  `mapstructure:"host"`
	Port      int     `mapstructure:"port"`
	URISchema string  `mapstructure:"uri-schema"`
	Version   string  `mapstructure:"version"`
	Log       logging `mapstructure:"log"`
}

// LoadConfig loads a configuration file from disk
func LoadConfig(path string) (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	config := viper.New()

	file := os.Getenv("SG_CONFIG")
	if file == "" {
		config.SetConfigName("config")
		config.AddConfigPath(path)
		config.AddConfigPath(fmt.Sprintf("%s/.config/7g", home))
	} else {
		base := filepath.Base(file)
		if strings.HasSuffix(base, "yaml") ||
			strings.HasSuffix(base, "json") ||
			strings.HasSuffix(base, "hcl") ||
			strings.HasSuffix(base, "toml") ||
			strings.HasSuffix(base, "conf") {
			// strip the file type for viper
			parts := strings.Split(filepath.Base(file), ".")
			base = strings.Join(parts[:len(parts)-1], ".")
		}
		config.SetConfigName(base)
		config.AddConfigPath(filepath.Dir(file))
	}

	err = config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config.SetEnvPrefix("SG")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	var c Config

	err = config.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	// initialize logging
	switch c.Log.Formatter {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	level, err := logrus.ParseLevel(c.Log.Level)
	if err == nil {
		logrus.SetLevel(level)
	}

	logrus.Debug("Configuration Location: ", config.ConfigFileUsed())

	return &c, nil
}

// ListenEndpoint builds the endpoint string (host + port)
func (c *Config) ListenEndpoint() string {
	if c.Port == 80 {
		return c.Host
	}
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// VersionedEndpoint builds the endpoint string (host + port + version)
func (c *Config) VersionedEndpoint(path string) string {
	return "/" + c.Version + path
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version)
func (c *Config) SchemaVersionedEndpoint(path string) string {
	if c.Port == 80 {
		return c.URISchema + c.Host + "/" + c.Version + path
	}
	return fmt.Sprintf("%s%s:%d/%s%s", c.URISchema, c.Host, c.Port, c.Version, path)
}
