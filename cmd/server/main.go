package main

import (
	"log"
	"net"

	"github.com/albimukti/assignment_3_albi/internal/cache"
	"github.com/albimukti/assignment_3_albi/internal/database"
	"github.com/albimukti/assignment_3_albi/internal/handler"
	"github.com/albimukti/assignment_3_albi/internal/proto"
	"github.com/albimukti/assignment_3_albi/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	connString := "postgres://username:password@localhost:5432/yourdbname?sslmode=disable"
	database.InitDB(connString)
	defer database.CloseDB()

	cache.InitCache()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterURLShortenerServer(grpcServer, &service.URLShortenerService{})
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	r := gin.Default()
	r.GET("/:short_url", handler.RedirectHandler)
	r.Run(":8080")
}
