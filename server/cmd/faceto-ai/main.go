package main

import (
	"errors"
	"flag"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/go-kratos/kratos/v2/config/env"
	"go.uber.org/zap/zapcore"

	"faceto-ai/internal/conf"
	"faceto-ai/internal/pkg/middleware"
	pkgLog "faceto-ai/internal/pkg/utils/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "go.uber.org/automaxprocs"
	_ "net/http/pprof"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	Env string

	id, _ = os.Hostname()

	logDir     = "logs"
	logFileExt = "info.log"
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&Env, "env", "", "env, eg: -env dev")
	flag.StringVar(&Name, "name", "faceto-ai", "name, eg: -name XXX")
}

func newApp(logger log.Logger, gs *grpc.Server, sl *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			sl,
		),
	)
}

func initLog() (log.Logger, func()) {
	if _, err := os.Stat(logDir); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(logDir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	logFileName := Name + "." + logFileExt
	infoLogFilePath := path.Join(logDir, logFileName)
	logFile, err := os.Create(infoLogFilePath)
	if err != nil {
		log.Fatal("failed to open error log file: ", err)
	}

	p, _ := filepath.Abs(logFile.Name())
	log.Infof("log file: %s", p)

	logger := pkgLog.InitFileLogger(logFile, zapcore.DebugLevel)
	return logger, func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}
}

func initStdLog() (log.Logger, func()) {
	logger := pkgLog.InitStdLogger(zapcore.DebugLevel)
	return logger, func() {
		if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
			panic(err)
		}
	}
}

func initProductStdLog() (log.Logger, func()) {
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	return logger, func() {
		if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
			panic(err)
		}
	}
}

func main() {
	flag.Parse()

	var (
		logger    log.Logger
		cleanFunc func()
	)
	if Env == "dev" {
		logger, cleanFunc = initStdLog()
	} else {
		logger, cleanFunc = initProductStdLog()
	}
	defer cleanFunc()

	logger = log.With(logger,
		"service.id", id,
		"service.version", Version,
		"trace.id", middleware.RequestID(),
	)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
			env.NewSource(),
		),
	)

	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if Env == "dev" {
		bc.Data.Database.Debug = true
		bc.Server.Debug = conf.ServerDebug_true
	}

	app, cleanup, err := wireApp(&bc, bc.Server, bc.Data, bc.ThirdApi, bc.Storage, bc.Livekit, bc.GcpCredentials, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	//go func() {
	//	netHttp.ListenAndServe("0.0.0.0:8899", nil)
	//}()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
