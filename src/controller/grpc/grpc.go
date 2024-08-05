package grpc

import (
	"context"
	"time"

	pbAsset "github.com/kowiste/boilerplate/doc/proto/asset"
	pbUser "github.com/kowiste/boilerplate/doc/proto/user"
	conf "github.com/kowiste/boilerplate/src/config"
	assetservice "github.com/kowiste/boilerplate/src/service/asset"
	userservice "github.com/kowiste/boilerplate/src/service/user"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"log"
	"net"

	"github.com/kowiste/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type GRPC struct {
	server       *grpc.Server
	serviceUser  *userservice.UserService
	serviceAsset *assetservice.AssetService
	// Embed the unimplemented server structs
	pbAsset.UnimplementedAssetServiceServer
	pbUser.UnimplementedUserServiceServer
	tracer trace.Tracer
}
type Option func(*GRPC) error

func New(opts ...Option) (g *GRPC) {
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
	g = &GRPC{
		server:       grpc.NewServer(),
		serviceUser:  sU,
		serviceAsset: sA,
	}
	g.applyOptions(opts...)
	return
}

// WithTracer sets the tracer for the API instance.
func WithTracer(tracer *trace.Tracer) Option {
	return func(a *GRPC) error {
		a.tracer = *tracer
		return nil
	}
}

func (g *GRPC) applyOptions(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(g); err != nil {
			return err
		}
	}
	return nil
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

func (a GRPC) TelemetryUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	ctx, span := a.tracer.Start(ctx, info.FullMethod)
	defer span.End()

	// Record the start time
	startTime := time.Now()

	// Handle the request
	resp, err = handler(ctx, req)

	// Record the end time
	endTime := time.Now()
	latency := endTime.Sub(startTime)
	st, _ := status.FromError(err)

	// Set span attributes
	span.SetAttributes(
		attribute.String("grpc.method", info.FullMethod),
		attribute.String("grpc.client_ip", getClientIP(ctx)),
		attribute.Int("grpc.status_code", int(st.Code())),
		attribute.String("grpc.user_agent", getUserAgent(ctx)),
		attribute.Float64("grpc.latency", latency.Seconds()),
	)

	return resp, err
}

// Helper function to get client IP from context (you might need to implement this)
func getClientIP(ctx context.Context) string {
	// Implement your logic to retrieve client IP from context metadata
	return "127.0.0.1"
}

// Helper function to get user agent from context (you might need to implement this)
func getUserAgent(ctx context.Context) string {
	// Implement your logic to retrieve user agent from context metadata
	return "grpc-client/1.0"
}
