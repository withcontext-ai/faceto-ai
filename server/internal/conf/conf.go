// Package conf
// ---------------------------------
// @file      : conf.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/9/20 18:33
// @desc      : file description
// ---------------------------------
package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func NewConf() *Bootstrap {

	configPath := "../../../configs/config.yaml"

	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
			env.NewSource(),
		),
	)

	if err := c.Load(); err != nil {

		panic(err)
	}

	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return &bc
}

func GetGcpCredBytes(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("gcp cred url empty")
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Google GCP Credentials err")
	}
	defer response.Body.Close()

	gcpData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Google GCP Credentials Resp Read err")
	}
	return gcpData, nil
}
