package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shortener/config"
	"shortener/grpc_domain"
	"shortener/internal/repo"
	"shortener/internal/usecase"
	"shortener/pkg/postgres"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cfg *config.Config) {
	f, err := os.OpenFile(cfg.PathLog+"/log.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open log file: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Println(cfg.DBURL)
	log.Println(cfg.StorageMode)
	var repository repo.Repository
	if cfg.StorageMode == "in-memory" {
		repository = repo.NewIM()
	} else if cfg.StorageMode == "postgresql" {
		if len(cfg.DBURL) == 0 {
			log.Fatalln("error db url")
		}
		pool, err := postgres.New(cfg.DBURL)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer pool.Close()
		repository = repo.NewPG(pool)
	} else {
		log.Fatalln("error storage mode")
	}

	uc := usecase.New(repository)

	lis, err := net.Listen("tcp", cfg.PortGRPC)
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}
	s := grpc.NewServer()

	grpc_domain.RegisterShortenerServer(s, uc)

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalln("failed to serve: ", err)
		}
	}()

	//Client connection
	conn, err := grpc.DialContext(
		context.Background(),
		cfg.PortGRPC,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	err = grpc_domain.RegisterShortenerHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    cfg.PortHTTP,
		Handler: gwmux,
	}

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
		if err := gwServer.Shutdown(shutdownCtx); err != nil {
			log.Fatalln("gwserver shutdown error: ", err)
		}

		serverStopCtx()
	}()

	log.Println("Serving gRPC-Gateway on port ", cfg.PortHTTP)
	if err = gwServer.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed: ", err)
			os.Exit(0)
		}
		log.Fatalln("failed to listen and serve: ", err)
	}
}
