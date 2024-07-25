package grpc

import (
	pbAsset "boiler/doc/proto/asset"
	pbUser "boiler/doc/proto/user"
	conf "boiler/src/config"
	assetservice "boiler/src/service/asset"
	userservice "boiler/src/service/user"

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

	//pbAsset.RegisterAssetServiceServer(a.server, a)
	pbUser.RegisterUserServiceServer(a.server, a)

	log.Println("Starting gRPC server on port " + c.GRPCPort)
	go func() {
		if err := a.server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	return nil
}
