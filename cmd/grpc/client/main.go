package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"grpc-comm/internal/grpc/domain"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	client := domain.NewResositoryServiceClient(conn)


	for i := range [100000]int{} {
		repostioryModel := domain.RepositoryRequest{
			Id:        int64(i),
			Name:      string("Grpc-to-go"),
			UserId:    12345,
			IsPrivate: true,
		}
		wg.Add(1)
		go retrieveData(repostioryModel, client)
		
	}
	wg.Wait()
}

var wg sync.WaitGroup

func retrieveData(repo domain.RepositoryRequest, client domain.ResositoryServiceClient) {
	defer wg.Done()
	if responseMessage, err := client.Add(context.Background(), &repo); err != nil {
		panic(fmt.Sprintf("Was not able to insert Record %v", err))
	} else {
		fmt.Println("Record Inserted..")
		fmt.Println(responseMessage)
		fmt.Println("=============================")
	}
}
