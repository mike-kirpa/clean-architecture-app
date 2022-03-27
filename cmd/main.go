package main

import (
	"clean-architecture-app/internal/composites"
	"clean-architecture-app/internal/config"
	"clean-architecture-app/pkg/logging"
	"context"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Port, cfg.MongoDB.Port, "", "", "", "")
	if err != nil {
		logger.Fatal("mongoDB composite failed")
	}

	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(*mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	authorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewAuthorComposite(*mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}
	bookComposite.Handler.Register(router)

	//start
}
