package factory

import (
	"time"
)

type ManualConfig struct {
	NFs map[string][]NF `yaml:"nfs,omitempty"` // Map of NF types to their configurations
}

type NF struct {
	Name          string        `yaml:"name"`
	Description   string        `yaml:"description,omitempty"`
	CreatedAt     time.Time     `yaml:"createdAt,omitempty"`
	UpdatedAt     time.Time     `yaml:"updatedAt,omitempty"`
	URL           string        `yaml:"url"`
	Port          int           `yaml:"port"`
	FiltersParams FiltersParams `yaml:"filtersParams,omitempty"`
	ServicesList  []string      `yaml:"servicesList"`
}

type FiltersParams struct {
	SupiRangeStart string `yaml:"supiRangeStart,omitempty"`
	SupiRangeEnd   string `yaml:"supiRangeEnd,omitempty"`
	GroupId        string `yaml:"groupId,omitempty"`
}
