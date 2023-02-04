package conf

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Language struct {
	Name           string   `yaml:"name"`
	Acros          []string `yaml:"acros"`
	Path           string   `yaml:"path"`
	TemplatePath   string   `yaml:"templatePath"`
	InitialCommand []string `yaml:"initialCommand"`
	EditorPath     string   `yaml:"editorPath"`
}

type Config struct {
	Languages                 []Language `yaml:"languages"`
	Vcs                       string     `yaml:"vcs"`
	InferLanguage             bool       `yaml:"inferLanguage"`
	GhKey                     string     `yaml:"ghKey"`
	DefaultEditorPath         string     `yaml:"defaultEditorPath"`
	DefaultVcsState           bool       `yaml:"defaultVcsState"`
	DefaultGithubVisibility   bool       `yaml:"defaultGithubVisibility"`
	DefaultCreateREADME       bool       `yaml:"defaultCreateREADME"`
	DefaultCreateRequirements bool       `yaml:"defaultCreateRequirements"`
	DisableExtensions         bool       `yaml:"disableExtensions"`
}

func GetConfig(yamlBytes []byte) (Config, error) {
	var conf Config = Config{}

	var err3 = yaml.Unmarshal(yamlBytes, &conf)

	if err3 != nil {
		return Config{}, err3
	}

	return validateData(conf)
}

// used to check whether values entered in configPath are valid
//
//	i.e Language.Path != "" ...
func validateData(config Config) (Config, error) {
	for _, v := range config.Languages {
		if v.Path == "" {
			return Config{}, fmt.Errorf("language path must not be empty. Please update config.yaml")
		} else if len(v.Acros) == 0 {
			return Config{}, fmt.Errorf("language acros must not be empty. Please update config.yaml")
		}
	}
	if config.DefaultEditorPath == "" {
		return Config{}, fmt.Errorf("DefaultEditorPath must not be empty. Please update config.yaml")
	}

	if config.GhKey == "" && config.Vcs == "github" {
		return Config{}, fmt.Errorf("ghKey must be set to use github version control system. Please update config.yaml")
	}

	return config, nil
}
