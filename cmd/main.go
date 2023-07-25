package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"freelance/admin_panel/api"
	"freelance/admin_panel/api/handlers"
	"freelance/admin_panel/config"
	"freelance/admin_panel/storage/mongodb"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	// credential := options.Credential{
	// 	Username:      cfg.MongoUser,
	// 	// Password:      cfg.MongoPassword,
	// 	// AuthMechanism: "SCRAM-SHA-256",
	// 	AuthSource:    cfg.MongoDatabase,
	// }

	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort)
	log.Println("MAIN TEST PRINT " + mongoString)
	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString)) // .SetAuth(credential)
	if err != nil {
		log.Fatal("error to connect to mongo database", logger.Error(err))
	}

	defer func(mongoConn *mongo.Client, ctx context.Context) {
		err := mongoConn.Disconnect(ctx)
		if err != nil {
			return
		}
	}(mongoConn, context.Background())

	if err := mongoConn.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Cannot connect to database error -> ", logger.Error(err))
	}
	connDB := mongoConn.Database(cfg.MongoDatabase)

	store := mongodb.NewStoragePg(connDB)

	logger.Any("Connected to MongoDB in ", logger.Any("Server: ", cfg.MongoHost))

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	h := handlers.NewHandler(cfg, log, store)

	r := api.SetUpRouter(h, cfg)

	r.Run(cfg.HTTPPort)

}
