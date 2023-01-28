package gapi

import (
	"context"
	"database/sql"

	db "github.com/redsubmarine/simplebank/db/sqlc"
	"github.com/redsubmarine/simplebank/pb"
)

func (server *Server) ListItems(ctx context.Context, req *pb.ListItemRequest) (*pb.ListItemResponse, error) {
	id := req.GetId()
	limit := req.GetLimit()
	if limit == 0 {
		limit = 10
	}

	arg := db.ListItemsParams{
		ID: sql.NullInt32{
			Int32: id,
			Valid: id > 0,
		},
		Limit: limit,
	}

	list, err := server.store.ListItems(ctx, arg)
	if err != nil {
		return nil, err
	}

	rsp := &pb.ListItemResponse{
		List: convertItems(list),
	}
	return rsp, nil
}

func convertItems(list []db.Item) []*pb.Item {
	items := make([]*pb.Item, len(list))
	for i, item := range list {
		items[i] = convertItem(item)
	}
	return items
}

func convertItem(item db.Item) *pb.Item {
	return &pb.Item{
		Id:   int32(item.ID),
		Name: item.Name,
	}
}
