package server

import (
	"context"
	"encoding/json"
	"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server/bapi"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())

	if err != nil {
		return nil, err
	}
	if body == nil || len(body) == 0 {
		return nil, errcode.ToRPCError(errcode.NotFound)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)

	if err != nil {
		return nil, errcode.ToRPCError(errcode.Fail)
	}

	return &tagList, nil
}
