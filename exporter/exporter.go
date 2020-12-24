package exporter

import (
	"github.com/ffddorf/unms-exporter/client"
	openapi "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Exporter struct {
	api *client.UNMSAPI
}

func New(token string) *Exporter {
	api := client.NewHTTPClient(strfmt.Default)

	client, ok := api.Transport.(*openapi.Runtime)
	if !ok {
		panic("Invalid openapi transport")
	}
	auth := openapi.APIKeyAuth("x-auth-token", "header", token)
	client.DefaultAuthentication = auth

	return &Exporter{
		api,
	}
}
