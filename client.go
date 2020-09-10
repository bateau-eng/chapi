// Package chapi wraps the generated swagger client in a nice client constructor.
package chapi

import (
	apiclient "github.com/bateau-eng/chapi/pkg/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

// NewClient creates a new apiclient that uses the provided key for all calls.
func NewClient(chKey string) *apiclient.ClubhouseAPI {
	c := apiclient.NewHTTPClientWithConfig(nil, &apiclient.TransportConfig{
		Host:     "api.clubhouse.io",
		Schemes:  []string{"https"},
		BasePath: apiclient.DefaultBasePath,
	})
	c.SetTransport(newAuthedTransport(c.Transport, chKey))
	return c
}

func newAuthedTransport(t runtime.ClientTransport, chKey string) runtime.ClientTransport {
	return &authedTransport{
		auth: client.APIKeyAuth("Clubhouse-Token", "header", chKey),
		t:    t,
	}
}

type authedTransport struct {
	auth runtime.ClientAuthInfoWriter
	t    runtime.ClientTransport
}

func (at *authedTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	op.AuthInfo = at.auth
	return at.t.Submit(op)
}
