package main

import (
	"os"

	provider "github.com/TsuyoshiUshio/go-plugin-sample/common"
	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
)

type AzureProvider struct {
	logger hclog.Logger
}

func (p *AzureProvider) Create(scheme *provider.Scheme) string {
	p.logger.Debug("Creating Azure Resources ...")
	return "123"
}

func (p *AzureProvider) Update(scheme *provider.Scheme) string {
	p.logger.Debug("Updating Azure Resources ...")
	return "123"
}

func (p *AzureProvider) Delete(scheme *provider.Scheme) string {
	p.logger.Debug("Deleting Azure Resources ...")
	return ""
}
func (p *AzureProvider) Get(id string) *provider.Scheme {
	p.logger.Debug("getting Azure Resources ...")
	scheme := &provider.Scheme{
		"azure vm",
	}
	return scheme
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "PROVIDER_PLUGIN",
	MagicCookieValue: "azure",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	azureProvider := &AzureProvider{
		logger: logger,
	}
	var pluginMap = map[string]plugin.Plugin{
		"azure": &provider.ProviderPlugin{Impl: azureProvider},
	}
	logger.Debug("Now Azure Provider Serving")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
