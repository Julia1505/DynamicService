package main

import (
	"UserSegmentationService/internal/handler"
	"UserSegmentationService/internal/repository"
	"UserSegmentationService/internal/service"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	storage, err := repository.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	dynService := service.NewService(storage)
	dynHandler := handler.NewHandler(dynService)
	router := dynHandler.InitRouter()

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
