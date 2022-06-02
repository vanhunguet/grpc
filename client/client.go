package main

import (
	"context"
	"fmt"
	pb "github.com/vanhunguet/grpc/gen/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	res, err := client.GetUser(context.Background(), &pb.UserRequest{Uuid: ""})

	fmt.Println(res)
}
