// Package data
// ---------------------------------
// @file      : data_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/9/15 16:19
// @desc      : file description
// ---------------------------------
package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/ent"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap/zapcore"
	"testing"
)

var (
	ctx    context.Context
	logger log.Logger
	dbCfg  *conf.Data
)

func TestMain(t *testing.M) {
	ctx = context.Background()
	logger = pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	dbCfg = &conf.Data{
		Database: &conf.Data_Database{
			Driver: "",
			Source: "",
			Debug:  true,
		},
	}
}

func NewMockClient() (*Data, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	driverOption := ent.Driver(sql.OpenDB("mysql", db))
	client := ent.NewClient(driverOption)

	if dbCfg.Database.Debug {
		client = client.Debug()
	}

	d := &Data{
		db: client,
	}
	return d, mock, nil
}

func TestLinkRepo_Count(t *testing.T) {
	dbData, _, err := NewMockClient()
	if err != nil {
		t.Errorf("TestNewData.NewMockClient err:%v", err)
	}

	// mock执行指定SQL语句时的返回结果
	//mock.ExpectBegin()
	//mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectCommit()

	//mock.ExpectExec("SELECT count(1) FROM link WHERE deleted_at is null").WillReturnResult(sqlmock.NewResult(10, 1))
	//fmt.Println("mock.ExpectExec")

	linkRepoClient := NewLinkRepo(dbData, logger)
	n, err := linkRepoClient.Count(ctx)
	if err != nil {
		t.Errorf("TestNewData.linkRepoClient.Count err:%v", err)
	}
	fmt.Println(n)

	link, err := linkRepoClient.GetLinkByName(ctx, "87ai-1t02")
	if err != nil {
		t.Errorf("TestNewData.linkRepoClient.GetLinkByName err:%v", err)
	}
	fmt.Println(link.RoomName)
}
