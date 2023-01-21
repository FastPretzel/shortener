package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"shortener/config"
	"shortener/grpc_domain"
	"shortener/internal/repo"
	"shortener/internal/usecase"
	"shortener/pkg/postgres"
	"syscall"
	"time"

	"google.golang.org/grpc"
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
	if cfg.DBURL == "" || cfg.InMemoryMode == true {
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

	lis, err := net.Listen("tcp", cfg.PortGRPC)
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}
	s := grpc.NewServer()

	grpc_domain.RegisterShortenerServer(s, uc)
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

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
}
