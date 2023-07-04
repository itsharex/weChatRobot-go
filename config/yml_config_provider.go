package config

import (
	"errors"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"path/filepath"
	"weChatRobot-go/models"
)

type ymlConfigProvider struct {
	filePath string
}

// NewFile returns a new Provider that reads the configuration from the given file.
func NewFile(filePath string) Provider {
	return &ymlConfigProvider{
		filePath: filePath,
	}
}

func (ycp *ymlConfigProvider) RetrieveConfig() (*models.ConfigSettings, error) {
	if ycp.filePath == "" {
		return nil, errors.New("config file not specified")
	}

	var config models.ConfigSettings
	absolutePath, err := filepath.Abs(ycp.filePath)
	if err != nil {
		return nil, err
	}

	k := koanf.New("::")
	_ = k.Load(file.Provider(absolutePath), yaml.Parser())
	err = k.Unmarshal("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}
