package grpc

import (
	pbAsset "github.com/kowiste/boilerplate/doc/proto/asset"
	pbUser "github.com/kowiste/boilerplate/doc/proto/user"
	conf "github.com/kowiste/boilerplate/src/config"
	assetservice "github.com/kowiste/boilerplate/src/service/asset"
	userservice "github.com/kowiste/boilerplate/src/service/user"

	"log"
	"net"

	"github.com/kowiste/config"
	"google.golang.org/grpc"
)

type GRPC struct {
	server       *grpc.Server
	serviceUser  *userservice.UserService
	serviceAsset *assetservice.AssetService
	// Embed the unimplemented server structs
	pbAsset.UnimplementedAssetServiceServer
	pbUser.UnimplementedUserServiceServer
}

func New() (g *GRPC) {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	sU, err := userservice.Get()
	if err != nil {
		return
	}
	sA, err := assetservice.Get()
	if err != nil {
		return
	}
	return &GRPC{
		server:       grpc.NewServer(),
		serviceUser:  sU,
		serviceAsset: sA,
	}
}
func (a *GRPC) Init() (err error) {
	c, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		return
	}
	lis, err := net.Listen("tcp", ":"+c.GRPCPort)
	if err != nil {
		return err
	}

	pbAsset.RegisterAssetServiceServer(a.server, a)
	pbUser.RegisterUserServiceServer(a.server, a)

	log.Println("Starting gRPC server on port " + c.GRPCPort)
	go func() {
		if err := a.server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	return nil
}
