package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Noah-Huppert/wgd/server/rpc"

	"github.com/Noah-Huppert/goconf"
	"github.com/Noah-Huppert/golog"
	//"github.com/vishvananda/netlink"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOpts "go.mongodb.org/mongo-driver/mongo/options"
	mongoReadpref "go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	grpcCredentials "google.golang.org/grpc/credentials"
)

// Config stores application configuration parameters.
type Config struct {
	// RPC configuration.
	RPC struct {
		// ListenAddr is the address on which the GRPC socket will listen.
		ListenAddr string `default:"127.0.0.1:6000" validate:"required"`

		// KeyFile is the encryption key file.
		KeyFile string `default:"../rpc/dev-server.key" validate:"required"`

		// CertificateFile is the encryption certificate file.
		CertificateFile string `default:"../rpc/dev-server.pem" validate:"required"`
	}

	// MongoDB holds MongoDB configuration.
	MongoDB struct {
		// URI is the address of the MongoDB database.
		URI string `default:"mongodb://127.0.0.1:27017" validate:"required"`

		// DatabaseName is the name of the MongoDB database.
		DatabaseName string `default:"dev_wgd" validate:"required"`
	}
}

// DB holds any database related context.
type DB struct {
	// Client is the MongoDB client.
	Client *mongo.Client

	// Database is a handle for the Mongo database.
	Database *mongo.Database

	// Users is a handle to the users Mongo collection.
	Users *mongo.Collection

	// Subnets is a handle to the subnets Mongo collection.
	Subnets *mongo.Collection
}

// RPCServer implements RPC interface.
type RPCServer struct {
	// Logger.
	Logger golog.Logger

	// Config.
	Config Config

	// DB.
	DB DB
}

// Listen on the configured port for GRPC requests.
func (s *RPCServer) Listen() error {
	// Load credentials
	creds, err := grpcCredentials.NewServerTLSFromFile(
		s.Config.RPC.CertificateFile,
		s.Config.RPC.KeyFile)
	if err != nil {
		return fmt.Errorf("failed to load server credentials: %s", err)
	}

	// Setup server
	rpcListen, err := net.Listen("tcp", s.Config.RPC.ListenAddr)
	if err != nil {
		return fmt.Errorf("error listening on \"%s\": %s",
			s.Config.RPC.ListenAddr, err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	rpc.RegisterRegistryServer(grpcServer, &Registry{
		Logger: s.Logger.GetChild("registry"),
		Config: s.Config,
	})

	s.Logger.Infof("Listening for GRPC requests on \"%s\"",
		s.Config.RPC.ListenAddr)
	grpcServer.Serve(rpcListen)

	return nil
}

// Implements the Registry RPC service.
type Registry struct {
	// Logger.
	Logger golog.Logger

	// Config.
	Config Config
}

func (r *Registry) HealthCheck(ctx context.Context, req *rpc.HealthCheckRequest) (*rpc.HealthCheckResponse, error) {
	return nil, nil
}

func (r *Registry) CreateInvitedUser(ctx context.Context, req *rpc.CreateInvitedUserRequest) (*rpc.CreateInvitedUserResponse, error) {
	return &rpc.CreateInvitedUserResponse{}, nil
}

func (r *Registry) ApproveInvitedUser(ctx context.Context, req *rpc.ApproveInvitedUserRequest) (*rpc.ApproveInvitedUserResponse, error) {
	return nil, nil
}

func (r *Registry) GetUsers(req *rpc.GetUsersRequest, resp_stream rpc.Registry_GetUsersServer) error {
	return nil
}

func (r *Registry) UpdateUser(ctx context.Context, req *rpc.UpdateUserRequest) (*rpc.UpdateUserResponse, error) {
	return &rpc.UpdateUserResponse{}, nil
}

func (r *Registry) DeleteUser(ctx context.Context, req *rpc.DeleteUserRequest) (*rpc.DeleteUserResponse, error) {
	return &rpc.DeleteUserResponse{}, nil
}

func (r *Registry) CreateSubnet(ctx context.Context, req *rpc.CreateSubnetRequest) (*rpc.CreateSubnetResponse, error) {
	return &rpc.CreateSubnetResponse{}, nil
}

func (r *Registry) GetSubnets(req *rpc.GetSubnetsRequest, resp_stream rpc.Registry_GetSubnetsServer) error {
	return nil
}

func (r *Registry) UpdateSubnetMeta(ctx context.Context, req *rpc.UpdateSubnetMetaRequest) (*rpc.UpdateSubnetMetaResponse, error) {
	return &rpc.UpdateSubnetMetaResponse{}, nil
}

func (r *Registry) AssignSubnetAddress(ctx context.Context, req *rpc.AssignSubnetAddressRequest) (*rpc.AssignSubnetAddressResponse, error) {
	return &rpc.AssignSubnetAddressResponse{}, nil
}

func (r *Registry) RemoveSubnetAddress(ctx context.Context, req *rpc.RemoveSubnetAddressRequest) (*rpc.RemoveSubnetAddressResponse, error) {
	return &rpc.RemoveSubnetAddressResponse{}, nil
}

func (r *Registry) DeleteSubnet(ctx context.Context, req *rpc.DeleteSubnetRequest) (*rpc.DeleteSubnetResponse, error) {
	return &rpc.DeleteSubnetResponse{}, nil
}

func main() {
	ctx := context.Background()

	// Logger
	logger := golog.NewLogger("wgd.server")
	logger.Debug("Starting")

	// Load configuration
	cfgLdr := goconf.NewLoader()
	cfgLdr.AddConfigPath("/etc/wgd/server.d/*")
	config := Config{}
	if err := cfgLdr.Load(&config); err != nil {
		logger.Fatalf("Failed to load configuration: %s", err)
	}

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(ctx,
		mongoOpts.Client().ApplyURI(config.MongoDB.URI))
	if err != nil {
		logger.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	if err := mongoClient.Ping(ctx, mongoReadpref.Primary()); err != nil {
		logger.Fatalf("Failed to ping MongoDB: %s", err)
	}

	mongoDatabase := mongoClient.Database(config.MongoDB.DatabaseName)

	db := DB{
		Client:   mongoClient,
		Database: mongoDatabase,
		Users:    mongoDatabase.Collection("users"),
		Subnets:  mongoDatabase.Collection("subnets"),
	}

	// Start GRPC server
	rpcServer := RPCServer{
		Logger: logger,
		Config: config,
		DB:     db,
	}
	if err := rpcServer.Listen(); err != nil {
		logger.Fatalf("Failed to run the RPC server "+
			"on \"%s\": %s", config.RPC.ListenAddr, err)
	}
}
