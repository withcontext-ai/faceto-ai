//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"faceto-ai/internal/biz"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data"
	"faceto-ai/internal/server"
	"faceto-ai/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(
	*conf.Bootstrap,
	*conf.Server,
	*conf.Data,
	*conf.ThirdApi,
	*conf.Storage,
	*conf.LiveKit,
	*conf.GcpCredentials,
	log.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
