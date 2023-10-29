package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app"
	"github.com/K-Kizuku/Somniosus_back/apps/account/config"
	"github.com/K-Kizuku/Somniosus_back/proto.pb/twitter/account"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func StartServer() {
	// DIセットアップ
	app, err := app.NewApp()
	if err != nil {
		panic(fmt.Errorf("wire initialize failed: %w", err))
	}
	defer app.DB.Close()
	// gRPCサーバー
	lis, err := net.Listen("tcp", ":"+config.AccountGrpcServerPort)
	if err != nil {
		panic(fmt.Errorf("network I/O error: %w", err))
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			app.Interceptor.UnaryAuthInterceptor(),
		),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)
	account.RegisterAccountServiceServer(s, app.AccountService)

	// HealthCheck
	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthSrv)
	healthSrv.SetServingStatus("account", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(s)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("serve error: %v", err)
		}
	}()
	defer s.GracefulStop()

	<-quit
}
