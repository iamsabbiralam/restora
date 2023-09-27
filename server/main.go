package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/iamsabbiralam/restora/server/storage/postgres"
	"github.com/iamsabbiralam/restora/utility"
	"github.com/iamsabbiralam/restora/utility/logging"

	"github.com/iamsabbiralam/restora/utility/middleware"
	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
	catC "github.com/iamsabbiralam/restora/server/core/categories"
	catS "github.com/iamsabbiralam/restora/server/service/categories"
)

var (
	svcName = "server"
	version = "development"
)

func main() {
	log.Printf("starting %s service", svcName)
	cfg, err := utility.NewConfig("env/config")
	if err != nil {
		log.Fatal(err)
	}
	
	logger := logging.NewLogger(cfg).WithFields(logrus.Fields{
		"service": svcName,
		"version": version,
	})
	
	dbString := utility.NewDBString(cfg)
	db, err := postgres.NewStorage(dbString, logger)
	if err != nil {
		logger.WithError(err).Error("unable to connect DB")
		return
	}
	
	if err := db.RunMigration(cfg.GetString("database.migrationDir")); err != nil {
		logger.WithError(err).Error("migration failed")
	}
	
	grpcServer, err := setupGRPCServer(db, cfg, logger)
	if err != nil {
		logger.WithError(err).Error("unable to setup grpc service")
		return
	}
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GetInt("server.port")))
	if err != nil {
		logger.WithError(err).Error("unable to listen port")
		return
	}
	
	log.Printf("server %s listening at: %+v", svcName, lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.WithError(err).Error("unable to serve the GRPC server")
	}

	log.Println("server stopped")
}

func setupGRPCServer(store *postgres.Storage, config *viper.Viper, logger *logrus.Entry) (*grpc.Server, error) {
	mw := middleware.New(
		config.GetString("runtime.environment"),
		logger,
		middleware.Config{},
	)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(mw),
	)

	coreCat := catC.New(store, logger)
	svcCat := catS.New(coreCat, logger)
	catG.RegisterCategoryServiceServer(grpcServer, svcCat)

	return grpcServer, nil
}
