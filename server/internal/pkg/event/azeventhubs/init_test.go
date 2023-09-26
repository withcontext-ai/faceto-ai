package azeventhubs

import (
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"

	"faceto-ai/internal/conf"
)

func newLog() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout))
}

func newConf() *conf.Bootstrap {
	configPath := "../../../../configs/"

	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
			env.NewSource(),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return &bc
}
