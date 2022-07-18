package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// Write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)

	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// Return response
	res := &logs.LogResponse{Result: "Logged"}
	return res, nil
}

func (app *Config) grpcListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for GRPC in %s", gRpcPort)
	}

	server := grpc.NewServer()
	logs.RegisterLogServiceServer(server, &LogServer{Models: app.Models})

	log.Printf("GRPC Server started on port %s", gRpcPort)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for GRPC in %s", gRpcPort)
	}
}
