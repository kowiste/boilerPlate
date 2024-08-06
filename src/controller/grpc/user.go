package grpc

import (
	"context"

	pbUser "github.com/kowiste/boilerplate/pkg/proto/user"
	"github.com/kowiste/boilerplate/src/model/user"
)

func (a *GRPC) GetAllUsers(ctx context.Context, req *pbUser.GetAllUsersRequest) (*pbUser.GetAllUsersResponse, error) {
	input := &user.FindUsersInput{
		Text: req.Name,
		Age:  int(req.Age),
	}
	users, err := a.serviceUser.Users(ctx,input)
	if err != nil {
		return nil, err
	}
	return &pbUser.GetAllUsersResponse{Users: users.ToGRPC()}, nil
}

func (a *GRPC) GetUserById(ctx context.Context, req *pbUser.GetByIdRequest) (*pbUser.GetUserByIdResponse, error) {

	user, err := a.serviceUser.UserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pbUser.GetUserByIdResponse{User: user.ToGRPC()}, nil
}
