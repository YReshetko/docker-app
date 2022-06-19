package config

import (
	"encoding/json"
	"os"
)

type ItemType string

const (
	STATIC  ItemType = "STATIC"
	DYNAMIC ItemType = "DYNAMIC"
)

type ConfigMap map[string]interface{}

type App struct {
	Workspaces    []Workspace   `json:"workspaces"`
	Services      []Service     `json:"services"`
	ServiceGroups ServiceGroups `json:"service_groups"`
}
type Workspace struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ServiceGroups map[string][]Service

type Service struct {
	Name            string           `json:"name"`
	ItemType        ItemType         `json:"item_type"`
	DockerConfig    *DockerConfig    `json:"docker_config,omitempty"`
	GitConfig       *GitConfig       `json:"git_config,omitempty"`
	LocalRepoConfig *LocalRepoConfig `json:"local_repo_config,omitempty"`
}
type DockerConfig struct {
	Image          string    `json:"image"`
	ContainerName  string    `json:"container_name"`
	ComposeSnippet string    `json:"compose_snippet"`
	ConfigMap      ConfigMap `json:"config_map"`
	//BuildCommand  string   `json:"build_command"`
}

type GitConfig struct {
	Repo string `json:"repo"`
}
type Location struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
type LocalRepoConfig struct {
	Locations []Location `json:"locations"`
}

func LoadAppConfig(path string) (App, error) {
	var app App
	err := loadConfig(path, &app)
	return app, err
}

func loadConfig(path string, config interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}
