package grpc

import (
	pbAsset "boiler/doc/proto/asset"
	pbUser "boiler/doc/proto/user"
	assetservice "boiler/src/service/asset"
	userservice "boiler/src/service/user"

	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPC struct {
	server       *grpc.Server
	serviceUser  *userservice.UserService
	serviceAsset *assetservice.AssetService
}

func New() (g *GRPC, err error) {
	sU, err := userservice.Get()
	sA, err := assetservice.Get()
	return &GRPC{
		server: grpc.NewServer(),
		serviceUser: sU,
		serviceAsset: sA,
	}, err
}
func (a *GRPC) Init() (err error) {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	pbAsset.RegisterAssetServiceServer(a.server, a)
	pbUser.RegisterUserServiceServer(a.server, a)

	log.Println("Starting gRPC server on port 50051")
	go func() {
		if err := a.server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	return nil
}
