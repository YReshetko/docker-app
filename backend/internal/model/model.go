package model

import "github.com/YReshetko/docker-app/backend/internal/config"

type Model struct {
	config config.App
}

func NewModel(config config.App) *Model {
	return &Model{
		config: config,
	}
}

func (m *Model) Services() config.ServiceGroups {
	return m.config.ServiceGroups
}
