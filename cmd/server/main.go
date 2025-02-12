package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mostafanoorpur/order-sample/internal/config"
	"github.com/mostafanoorpur/order-sample/internal/httputil"
	"github.com/mostafanoorpur/order-sample/internal/order"
	"github.com/mostafanoorpur/order-sample/internal/order/presentation/http"
	orderRepo "github.com/mostafanoorpur/order-sample/internal/order/repo/postgres"
	"github.com/sirupsen/logrus"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	httpPkg "net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false,
	})

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	confPath := parseArgs()
	if strings.TrimSpace(confPath) == "" {
		panic("you've to declare config file using -c command. eg: -c=config.yml")
	}
	initConfig(confPath)

	cfg := config.GetConfig()

	postgresDSN := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDbName,
	)
	db, err := gorm.Open(gormPostgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("cannot connect to database")
	}

	e := echo.New()

	e.Validator = httputil.NewValidator(validator.New())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api/v1")

	orderService := order.NewOrderService(orderRepo.NewOrderPostgresRepository(db))

	userHttpController := http.NewOrderHttpController(orderService)
	userHttpController.RegisterRoutes(apiGroup)

	// Start server
	go func() {
		err = e.Start(":" + cfg.HttpPort)
		if err != nil && err != httpPkg.ErrServerClosed {
			logger.Fatal("shutting down the server", err)
		}
	}()

	// wait for `Ctrl+c` or docker stop/restart signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = e.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}

func parseArgs() string {
	var confPath string

	flag.StringVar(&confPath, "c", "./config.yaml", "config file for api")
	flag.Parse()

	return confPath
}

func initConfig(confPath string) {
	confPath, err := filepath.Abs(confPath)
	if err != nil {
		logrus.Fatal(err)
	}

	config.Init(confPath)
}
