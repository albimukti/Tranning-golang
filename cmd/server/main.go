package main

import (
	"log"
	"net"

	"github.com/albimukti/assignment_3_albi/internal/cache"
	"github.com/albimukti/assignment_3_albi/internal/database"
	"github.com/albimukti/assignment_3_albi/internal/handler"
	proto "github.com/albimukti/assignment_3_albi/internal/proto/proto"
	"github.com/albimukti/assignment_3_albi/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	connString := "postgres://postgres:albi123@localhost:5432/DB_tranning?sslmode=disable"
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
	r.POST("/short_url", handler.CreateShortURLHandler)
	r.Run(":8080")
}
