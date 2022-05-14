package impl

import (
	"context"
	"grpc-comm/internal/grpc/domain"
	"log"
	"strconv"
)

//RepositoryServerGrpcImpl is an impelmentation of RepositoryService Grpc Service
type RepositoryServiceGrpcImpl struct {
}

//NewRepositoryServiceGrpcImpl returns the pointer to the struct
func NewRepositoryServiceGrpcImpl() *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}

//Add function implentation of grpc service
func (serviceImpl *RepositoryServiceGrpcImpl) Add(ctx context.Context, req *domain.RepositoryRequest) (*domain.AddRepositoryResponse, error) {
	log.Println("Recieved request for adding repository with id", strconv.FormatInt(req.Id, 10))

	log.Println("Repository added to the storage")

	return &domain.AddRepositoryResponse{
		AddedRepository: req,
		Error:           nil,
	}, nil
}
