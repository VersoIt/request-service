package main

import (
	"RequestService/config"
	requestuc "RequestService/internal/app/usecase/request"
	requestservice "RequestService/internal/domain/service/request"
	"RequestService/internal/infrastructure/kafka"
	"RequestService/internal/infrastructure/pg"
	"RequestService/internal/infrastructure/pg/migrator"
	requestrepo "RequestService/internal/infrastructure/pg/repository/request"
	"RequestService/internal/server"
	"RequestService/internal/transport/http/request"
	"RequestService/pkg/validator"
	"context"
	trmgr "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	cfg := config.MustGet()

	db, err := pg.New(cfg)
	if err != nil {
		logrus.Panic(err)
	}

	pgMigrator := migrator.New(db)

	err = pgMigrator.Migrate()
	if err != nil {
		logrus.Panic(err)
	}

	kafkaProducer, err := kafka.New(cfg)
	if err != nil {
		logrus.Panic(err)
	}

	txManager := manager.Must(trmgr.NewDefaultFactory(db))
	requestRepo := requestrepo.New(db, trmgr.DefaultCtxGetter)
	requestService := requestservice.New(requestRepo)
	requestUC := requestuc.New(requestRepo, requestService, txManager, kafkaProducer, cfg)
	handler := request.NewHandler(requestUC)

	e := echo.New()
	e.Validator = validator.New()

	srv := server.New(cfg, e, handler)

	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	go func() {
		defer wg.Done()

		err := srv.Run(ctx)
		if err != nil {
			log.Error(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	cancel()

	wg.Wait()
}
