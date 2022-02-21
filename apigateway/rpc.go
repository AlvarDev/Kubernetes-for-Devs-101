package main

import (
	pb "apigatewayservice/pb"
	"context"
)

func (fe *frontendServer) getNewId(ctx context.Context, req *pb.GetNewIdRequest) (*pb.GetNewIdResponse, error) {
	return pb.NewReminderServiceClient(fe.getNewIdSvcConn).
		GetNewId(ctx, req)
}
