package config

import (
	"html/template"
)

// appConfig holds the application configuration.
type AppConfig struct {
	templateCache map[string]*template.Template
	isUsingCache  bool
	// infoLogger    *log.Logger
}

// NewAppConfig initializes the application configuration.
func NewAppConfig(tc map[string]*template.Template, useCache bool) *AppConfig {
	return &AppConfig{
		templateCache: tc,
		isUsingCache:  useCache,
	}
}

func (a *AppConfig) GetTemplateCache() map[string]*template.Template {
	return a.templateCache
}

func (a *AppConfig) GetIsUsingCache() bool {
	return a.isUsingCache
}
