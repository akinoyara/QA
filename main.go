package main

import (
	"context"
	"fmt"
	apihttp2 "restAPI/internal/apihttp"
	"restAPI/internal/storage/postgres"
	"restAPI/internal/storage/taskStorage"
)

func main() {

	ctx := context.Background()
	conn, err := postgres.CreateConnection(ctx)
	todoList := taskStorage.NewList(ctx, conn)
	httpHandlers := apihttp2.NewHTTPHandlers(todoList)
	httpServer := apihttp2.NewHTTPServer(httpHandlers)
	if err != nil {
		panic(err)
	}
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start HTTP server", err)
	}
}
