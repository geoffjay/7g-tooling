package context

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	bin "github.com/geoffjay/7g-tooling/data/init"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type database struct {
	File        string `mapstructure:"file"`
	Flush       bool   `mapstructure:"flush"`
	AutoMigrate bool   `mapstructure:"auto-migrate"`
}

type logging struct {
	Formatter string `mapstructure:"formatter"`
	Level     string `mapstructure:"level"`
}

// Configuration for the remote application host
type remote struct {
	Address string `mapstructure:"address"`
}

// Config defines the application configuration
type Config struct {
	Host      string   `mapstructure:"host"`
	Port      int      `mapstructure:"port"`
	URISchema string   `mapstructure:"uri-schema"`
	Version   string   `mapstructure:"version"`
	APIKey    string   `mapstructure:"api-key"`
	Database  database `mapstructure:"database"`
	Remote    remote   `mapstructure:"remote"`
	Log       logging  `mapstructure:"log"`
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

	// FIXME: hack to not have to pass config down as props to processor
	if err := os.Setenv("SG_API_KEY", c.APIKey); err != nil {
		logrus.Error(err)
	}
	if err := os.Setenv("SG_REMOTE_ADDRESS", c.Remote.Address); err != nil {
		logrus.Error(err)
	}
	logrus.Debug("Configuration Location: ", config.ConfigFileUsed())

	return &c, nil
}

func (c *Config) Write() {
	configPath := os.Getenv("SG_CONFIG")
	f, err := os.Create(configPath)
	if err != nil {
		logrus.Fatalf("Failed to open output file: %s", err)
	}
	t := template.Must(template.New("config").Parse(bin.GetAsset("config.yaml.tmpl")))
	err = t.Execute(f, c)
	if err != nil {
		logrus.Fatalf("Template generation failed: %s", err)
	}
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
