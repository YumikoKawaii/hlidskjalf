package api_key

import (
	"elysium.com/applications/utils"
	"gopkg.in/yaml.v2"
)

type ApiConfig struct {
	Route string   `yaml:"route"`
	Keys  []string `yaml:"keys"`
}

type ConfigMap struct {
	Configs []ApiConfig `yaml:"configs"`
}

type Resolver interface {
	Verify(route, key string) bool
}

type resolverImpl struct {
	keyMap map[string]map[string]bool
}

func NewResolver(path string) Resolver {

	data, err := utils.ReadFile(path)
	if err != nil {
		panic(err)
	}

	cfgMap := new(ConfigMap)
	if err := yaml.Unmarshal(data, cfgMap); err != nil {
		panic(err)
	}

	keyMap := make(map[string]map[string]bool)

	for _, cfg := range cfgMap.Configs {
		apiKeyMap := make(map[string]bool)
		for _, key := range cfg.Keys {
			apiKeyMap[key] = true
		}
		keyMap[cfg.Route] = apiKeyMap
	}

	return &resolverImpl{
		keyMap: keyMap,
	}
}

func (r *resolverImpl) Verify(route, key string) bool {
	return r.keyMap[route][key]
}
