package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lb/pb"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("Expected one argument. Got %d\n", len(os.Args)-1))
	}
	numOfRequests, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Expected an integer.")
	}
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Cannot estabilish a connection with %s\n", ":8080"))
	}
	i := 0
	for i < numOfRequests {
		n := rand.Int63n(10000)
		client := pb.NewMathClient(conn)
		result, err := client.GetFactors(context.Background(), &pb.IntValue{Value: n})
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else {
			log.Printf("%d -> %v\n", n, result)
		}
		i++
	}
}
