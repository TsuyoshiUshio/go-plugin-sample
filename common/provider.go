package provider

import (
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

type ProviderRPC struct{ client *rpc.Client }

func (p *ProviderRPC) Create(scheme *Scheme) string {
	var resp string
	err := p.client.Call("Plugin.Create", scheme, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
func (p *ProviderRPC) Update(scheme *Scheme) string {
	var resp string
	err := p.client.Call("Plugin.Update", scheme, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
func (p *ProviderRPC) Delete(scheme *Scheme) string {
	var resp string
	err := p.client.Call("Plugin.Delete", scheme, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
func (p *ProviderRPC) Get(id string) *Scheme {
	var resp *Scheme
	err := p.client.Call("Plugin.Get", id, resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type ProviderRPCServer struct {
	Impl Provider
}

func (s *ProviderRPCServer) Create(args *Scheme, resp *string) error {
	*resp = s.Impl.Create(args)
	return nil
}

type Scheme struct {
	Default interface{}
}

type Provider interface {
	Create(scheme *Scheme) string
	Update(scheme *Scheme) string
	Delete(scheme *Scheme) string
	Get(id string) *Scheme
}

type ProviderPlugin struct {
	Impl Provider
}

// Server
func (p *ProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ProviderRPCServer{Impl: p.Impl}, nil
}

// Client
func (ProviderPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ProviderRPC{client: c}, nil
}
