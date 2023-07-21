package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/iamsabbiralam/restora/usermgm/storage/postgres"
	utility "github.com/iamsabbiralam/restora/utility"
	"github.com/iamsabbiralam/restora/utility/logging"
	"github.com/iamsabbiralam/restora/utility/middleware"

	userG "github.com/iamsabbiralam/restora/proto/v1/usermgm/user"
	userC "github.com/iamsabbiralam/restora/usermgm/core/user"
	userS "github.com/iamsabbiralam/restora/usermgm/service/user"

	authG "github.com/iamsabbiralam/restora/proto/v1/usermgm/auth"
	authC "github.com/iamsabbiralam/restora/usermgm/core/auth"
	authS "github.com/iamsabbiralam/restora/usermgm/service/auth"
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

	log.Println("server stopped")
}

func setupGRPCServer(store *postgres.Storage, config *viper.Viper, logger *logrus.Entry) (*grpc.Server, error) {
	//	Authentication
	/* userID from client */
	// 	Authorization

	mw := middleware.New(
		config.GetString("runtime.environment"),
		logger,
		middleware.Config{},
		// iMD.UnaryServerInterceptor(),
		// pv.UnaryServerInterceptor(),
	)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(mw),
	)

	coreUsr := userC.New(store, logger)
	userSvc := userS.New(coreUsr, logger)
	userG.RegisterUserServiceServer(grpcServer, userSvc)

	coreAuth := authC.New(store, logger)
	svcAuth := authS.New(coreAuth, coreUsr, logger)
	authG.RegisterLoginServiceServer(grpcServer, svcAuth)

	return grpcServer, nil
}
