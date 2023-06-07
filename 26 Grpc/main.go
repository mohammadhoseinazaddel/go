package main

import (
	"TopLearn/Grpc/cmd"
	"TopLearn/Grpc/grpcserver"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	"google.golang.org/grpc/grpclog"
)

func main() {
	op := flag.String("op", "c", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		runGrpcServer()
	case "c":
		runGrpcClient()
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
	peopleServer, err := grpcserver.NewGrpcServer("root:1234@/people")
	if err != nil {
		log.Fatalln(err)
	}
	cmd.RegisterPersonServiceServer(server, peopleServer)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

func runGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := cmd.NewPersonServiceClient(conn)
	input := ""
	fmt.Println("All People? (y/n)")
	fmt.Scanln(&input)
	if strings.EqualFold(input, "y") {
		people, err := client.GetPeople(context.Background(), &cmd.Request{})
		if err != nil {
			log.Fatalln(err)
		}

		for {
			preson, err := people.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(preson)
		}
		return
	}

	fmt.Println("name?")
	fmt.Scanln(&input)

	person, err := client.GetPerson(context.Background(), &cmd.Request{Name: input})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(*person)
}
