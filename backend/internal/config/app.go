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
	ItemType        string          `json:"item_type"`
	DockerConfig    DockerConfig    `json:"docker_config"`
	GitConfig       GitConfig       `json:"git_config"`
	LocalRepoConfig LocalRepoConfig `json:"local_repo_config"`
}
type DockerConfig struct {
	ImageTag      string `json:"image_tag"`
	ContainerName string `json:"container_name"`
	BuildCommand  string `json:"build_command"`
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
