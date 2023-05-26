//auto generated by go-cli

package config

import (
	_ "embed"
	"github.com/candbright/go-core/config"
)

//go:embed application.yaml
var AppConfigData []byte

// AppConfig is global config parsed from application.yaml file
var AppConfig *config.Config

func init() {
	cfg, err := config.Parse(AppConfigData, config.YAML)
	if err != nil {
		panic(err)
	}
	AppConfig = cfg
}
