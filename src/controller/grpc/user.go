package grpc

import (
	pbUser "boiler/doc/proto/user"
	"context"
)

func (a *GRPC) GetAllUsers(ctx context.Context, req *pbUser.GetAllUsersRequest) (*pbUser.GetAllUsersResponse, error) {
	input := &userapi.FindUsersInput{
		Text: req.Name,
		Age:  int(req.Age),
	}
	users, err := a.serviceUser.Users(input)
	if err != nil {
		return nil, err
	}
	return &pbUser.GetAllUsersResponse{Users: users}, nil
}

func (a *GRPC) GetUserById(ctx context.Context, req *pbUser.GetByIdRequest) (*pbUser.GetUserByIdResponse, error) {
	user := a.serviceUser.GetUser()
	user.ID=req.Id
	user, err := a.serviceUser.UserByID(ctx)
	if err != nil {
		return nil, err
	}
	return &pbUser.GetUserByIdResponse{User: user}, nil
}
