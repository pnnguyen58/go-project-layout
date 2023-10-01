package controllers

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"

	"github.com/pnnguyen58/go-project-layout/configs"
	protogen "github.com/pnnguyen58/go-project-layout/pkg/proto_generated"
)

func ListenAndServe(ctx context.Context, logger *zap.Logger, loanHandler *Loan) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.C.Server.GRPCPort))
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Create a gRPC server instance
	grpcServer := grpc.NewServer()
	// Register our service with the gRPC server
	protogen.RegisterLoanServiceServer(grpcServer, loanHandler)
	// TODO: register more servers here

	// Serve gRPC server
	logger.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%v", configs.C.Server.GRPCPort))
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	maxMsgSize := 1024 * 1024 * 20
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("0.0.0.0:%v", configs.C.Server.GRPCPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	gwMux := runtime.NewServeMux()
	// Register service handlers
	err = protogen.RegisterLoanServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// TODO: register more service handlers here

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", configs.C.Server.HTTPPort),
		Handler: gwMux,
	}

	logger.Info(fmt.Sprintf("Serving gRPC-Gateway on port %v", configs.C.Server.HTTPPort))
	go func() {
		if err = gwServer.ListenAndServe(); err != nil {
			logger.Fatal(err.Error())
		}
	}()
	// Wait for a signal to shut down the server
	<-ctx.Done()

	// Gracefully stop the server
	grpcServer.GracefulStop()
}
