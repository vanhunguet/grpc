package main

import (
	"context"
	"fmt"
	pb "github.com/vanhunguet/grpc/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	//conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	res, err := client.GetUser(context.Background(), &pb.UserRequest{Uuid: 1})

	fmt.Println(res, 111)
}
