package main

import (
	"log"
	"net"

	"grpc-comm/internal/grpc/domain"
	"grpc-comm/internal/grpc/impl"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":7000")
	// netListener := getNetListener(7000)

	if err != nil {
		log.Fatal("unable to connect tcp", err.Error())
	}

	grpcServer := grpc.NewServer()

	repositoryServiceIml := impl.NewRepositoryServiceGrpcImpl()

	domain.RegisterResositoryServiceServer(grpcServer, repositoryServiceIml)

	//start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("unable to connect to grpc")
	}

}

// func getNetListener(port uint) net.Listener {
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 		panic(fmt.Sprintf("failed to listen: %v", err))
// 	}

// 	return lis
// }
