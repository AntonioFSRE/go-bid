package main

import (
	"fmt"
	"os"

	"github.com/AntonioFSRE/go-bid/internal/config"
	"github.com/AntonioFSRE/go-bid/internal/server"
	"github.com/AntonioFSRE/go-bid/pkg/logger"
	"github.com/AntonioFSRE/go-bid/pkg/store/postgres"
	"github.com/AntonioFSRE/go-bid/pkg/store/redis"
)

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs"
	}

	configName := os.Getenv("CONFIG_NAME")
	if configName == "" {
		configName = "config.local"
	}

	cfg, err := config.LoadConfig(configPath, configName)
	if err != nil {
		fmt.Printf("fatal error config file: %v\n", err)
		os.Exit(1)
	}

	log := logger.New()
	log.Init(cfg.Server.Debug, cfg.Logger.Level)

	dbConfig := postgres.NewConfig(
		cfg.DB.Driver,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSL,
	)
	db, err := postgres.NewClient(dbConfig)
	if err != nil {
		log.Fatalf("no db connection: %v", err)
	}
	defer db.Close()

	rdb := redis.New(&redis.Config{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}, log)
	if err := rdb.Open(); err != nil {
		log.Fatalf("no redis connection: %v", err)
	}
	defer rdb.Close()

	s := server.New(cfg, db, rdb, log)
	if err := s.Run(); err != nil {
		log.Panicf("this server is not running: %v", err)
	}
}
