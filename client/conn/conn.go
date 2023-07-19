package conn

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/iamsabbiralam/restora/utility/middleware"
)

type Conn struct {
	Server, Urm *grpc.ClientConn
}

func NewConns(logger *logrus.Entry, cfg *viper.Viper) *Conn {
	log.Printf("starting to dialing hrm service port: %s", cfg.GetString("services.hrmURL"))
	opts := getGRPCOpts(cfg)
	server, err := grpc.Dial(cfg.GetString("services.hrmURL"), opts...)
	if err != nil {
		logger.WithError(err).Fatal("failed to dial hrm service")
	}

	log.Printf("starting to dialing usermgm service port: %s", cfg.GetString("services.usermgmURL"))
	urm, err := grpc.Dial(cfg.GetString("services.usermgmURL"), opts...)
	if err != nil {
		logger.WithError(err).Fatal("failed to dial usermgm service")
	}

	return &Conn{
		Server: server,
		Urm:    urm,
	}
}

func (co *Conn) Close() {
	co.Server.Close()
	co.Urm.Close()
}

func getGRPCOpts(cnf *viper.Viper) []grpc.DialOption {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	opts = append(opts,
		grpc.WithBlock(), grpc.WithChainUnaryInterceptor(
			middleware.AuthForwarder(),
		))
	return opts
}
