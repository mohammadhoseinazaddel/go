package main

import (
	"TopLearn/Grpc/cmd"
	"TopLearn/Grpc/grpcserver"
	"flag"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	"google.golang.org/grpc/grpclog"
)

func main() {
	op := flag.String("op", "s", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		runGrpcServer()
	case "c":
		// call Client method
	}

}

func runGrpcServer() {
	grpclog.Println("Starting Server...")
	lis, err := net.Listen("tcp", ":8282")

	if err != nil {
		log.Fatalln("Failed to listen ", err)
	}
	grpclog.Println("listening on 127.0.0.1:8282 ")

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	peopleServer, err := grpcserver.NewGrpcServer("user:1234@/people")
	if err != nil {
		log.Fatalln(err)
	}
	cmd.RegisterPersonServiceServer(server, peopleServer)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
