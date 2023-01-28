package gapi

import (
	"context"

	"github.com/redsubmarine/simplebank/pb"
)

func (server *Server) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	arg := req.GetName()

	item, err := server.store.CreateItem(ctx, arg)
	if err != nil {
		return nil, err
	}
	rsp := &pb.CreateItemResponse{
		Item: &pb.Item{
			Id:   int32(item.ID),
			Name: item.Name,
		},
	}
	return rsp, nil
}
