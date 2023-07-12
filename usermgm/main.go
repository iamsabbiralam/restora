package main

import (
	"fmt"
	"log"
	"net"

	"github/iamsabbiralam/usermgm/storage/postgres"
	"github/iamsabbiralam/utility"
	"github/iamsabbiralam/utility/logging"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	svcName = "usermgm"
	version = "development"
)

func main() {
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
		logger.WithError(err).Error("unable to run DB migrations")
		return
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

	log.Println("server stoped")
}

func setupGRPCServer(store *postgres.Storage, config *viper.Viper, logger *logrus.Entry) (*grpc.Server, error) {
	//	Authentication
	/* userID from client */
	// 	Authorization
	

	grpcServer := grpc.NewServer(
	)

	return grpcServer, nil
}
