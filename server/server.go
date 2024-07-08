package main

import (
	"log"
	"net"
	ai "github.com/Salikhov079/military-ai/genprotos/ai"
	"github.com/Salikhov079/military-ai/service"
	postgres "github.com/Salikhov079/military-ai/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	ai.RegisterAiServiceServer(s, service.NewAIService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
