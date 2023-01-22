package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"shortener/config"
	"shortener/grpc_domain"
	"shortener/internal/repo"
	"shortener/internal/usecase"
	"shortener/pkg/postgres"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cfg *config.Config) {
	f, err := os.OpenFile(cfg.PathLog, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open log file: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)

	var repository repo.Repository
	if cfg.StorageMode == "in-memory" || cfg.DBURL == "" {
		repository = repo.NewIM()
	} else {
		pool, err := postgres.New(cfg.DBURL)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer pool.Close()
		repository = repo.NewPG(pool)
	}

	uc := usecase.New(repository)

	lis, err := net.Listen("tcp", cfg.Addr+cfg.PortGRPC)
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}
	s := grpc.NewServer()

	grpc_domain.RegisterShortenerServer(s, uc)
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	//Graceful	Shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		s.GracefulStop()

		serverStopCtx()
	}()

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
	//Client connection
	conn, err := grpc.DialContext(
		context.Background(),
		cfg.Addr+cfg.PortGRPC,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	gwmux := runtime.NewServeMux()
	err = grpc_domain.RegisterShortenerHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    cfg.Addr + cfg.PortHTTP,
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://", cfg.Addr, cfg.PortHTTP)
	if err = gwServer.ListenAndServe(); err != nil {
		log.Fatalln("failed to listen and serve: ", err)
	}
}
