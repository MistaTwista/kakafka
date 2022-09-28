package application

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

type Profile struct {
	Name             string   `yaml:"name" json:"name"`
	BootstrapServers []string `yaml:"bootstrapServers" json:"bootstrapServers"`
}

type Config struct {
	Version  string    `yaml:"version" json:"version"`
	Profiles []Profile `yaml:"profiles" json:"profiles"`
}

func loadConfig(ctx context.Context) (*Config, error) {
	runtime.LogDebug(ctx, "getting user config dir")

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	runtime.LogDebugf(ctx, "user config dir: %s", userConfigDir)

	runtime.LogDebug(ctx, "getting app config dir")
	appConfigDir := filepath.Join(userConfigDir, "kakafka")
	runtime.LogDebugf(ctx, "app config dir: %s", appConfigDir)

	runtime.LogDebug(ctx, "checking user config dir for existing")

	var needCreateDir bool

	switch stat, err := os.Stat(appConfigDir); {
	case err != nil:
		if os.IsNotExist(err) {
			needCreateDir = true
			break
		}

		return nil, err
	case !stat.IsDir():
		return nil, fmt.Errorf("%s is not directory", appConfigDir)
	}

	var needCreateConfig bool

	if needCreateDir {
		runtime.LogDebug(ctx, "creating user config dir")

		if err := os.Mkdir(appConfigDir, 0700); err != nil {
			return nil, errors.New("failed to create directory")
		}

		needCreateConfig = true
	}

	configFileName := filepath.Join(appConfigDir, "config.yaml")

	runtime.LogDebugf(ctx, "config file name: %s", configFileName)
	runtime.LogDebug(ctx, "checking user config file for existing")

	switch stat, err := os.Stat(configFileName); {
	case err != nil:
		if os.IsNotExist(err) {
			needCreateConfig = true
			break
		}

		return nil, err
	case stat.IsDir():
		return nil, fmt.Errorf("%s is directory", configFileName)
	}

	const emptyConfig = "---\nversion: \"0\"\n"

	if needCreateConfig {
		runtime.LogDebug(ctx, "creating user config file")

		if err := ioutil.WriteFile(configFileName, []byte(emptyConfig), 0600); err != nil {
			return nil, err
		}
	}

	runtime.LogDebug(ctx, "reading user config file")

	buf, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	runtime.LogDebug(ctx, "unmarshaling user config")

	var cfg Config
	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
