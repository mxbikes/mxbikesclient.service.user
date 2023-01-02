package main

import (
	"net"
	"os"

	"github.com/mxbikes/mxbikesclient.service.user/handlers"
	protobuffer "github.com/mxbikes/protobuf/user"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logLevel = GetEnv("LOG_LEVEL", "info")
	port     = GetEnv("PORT", ":4090")
)

func main() {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}

	//userHandler := handlers.NewUserHandler(*logger)
	//userHandler.GetUserByID("auth0|63b2dff9e834e550f0e50e66")
	/* Server */
	// Create a tcp listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "SERVICE.USER"}).Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	protobuffer.RegisterUserServiceServer(grpcServer, handlers.NewUserHandler(*logger))
	reflection.Register(grpcServer)

	// Start grpc server on listener
	logger.WithFields(logrus.Fields{"prefix": "SERVICE.USER"}).Infof("is listening on Grpc PORT: {%v}", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		logger.WithFields(logrus.Fields{"prefix": "SERVICE.USER"}).Fatalf("failed to serve: %v", err)
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
