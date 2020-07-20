package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Noah-Huppert/wgd/server/rpc"

	"github.com/Noah-Huppert/goconf"
	"github.com/Noah-Huppert/golog"
	//"github.com/vishvananda/netlink"
	"google.golang.org/grpc"
)

// Application config.
type Config struct {
	// RPCAddr is the address on which the GRPC socket will run.
	RPCAddr string `default:"127.0.0.1:6000" validate:"required"`
}

// Implements RPC interface.
type RPCServer struct {
	// Logger.
	Logger golog.Logger

	// Config.
	Config Config
}

// Listen on the configured port for GRPC requests.
func (s *RPCServer) Listen() error {
	rpcListen, err := net.Listen("tcp", s.Config.RPCAddr)
	if err != nil {
		return fmt.Errorf("error listening on \"%s\": %s",
			s.Config.RPCAddr, err)
	}

	grpcServer := grpc.NewServer()
	rpc.RegisterRegistryServer(grpcServer, &Registry{
		Logger: s.Logger.GetChild("registry"),
		Config: s.Config,
	})

	s.Logger.Infof("Listening for GRPC requests on \"%s\"",
		s.Config.RPCAddr)
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

// Create a new user and send them an invite code so they can get access to their account.
func (r *Registry) CreateInvitedUser(ctx context.Context, req *rpc.CreateInvitedUserRequest) (*rpc.CreateInvitedUserResponse, error) {
	return &rpc.CreateInvitedUserResponse{}, nil
}

// Retrieve details of users.
func (r *Registry) GetUsers(req *rpc.GetUsersRequest, resp_stream rpc.Registry_GetUsersServer) error {
	return nil
}

// Update a user's state. A user will be allowed to update only their state. Admins can
// update any user's state. The user argument should include the id field plus any fields
// which should be updated.
func (r *Registry) UpdateUser(ctx context.Context, req *rpc.UpdateUserRequest) (*rpc.UpdateUserResponse, error) {
	return &rpc.UpdateUserResponse{}, nil
}

// Delete a user and associated resources.
func (r *Registry) DeleteUser(ctx context.Context, req *rpc.DeleteUserRequest) (*rpc.DeleteUserResponse, error) {
	return &rpc.DeleteUserResponse{}, nil
}

// Create a new subnet.
func (r *Registry) CreateSubnet(ctx context.Context, req *rpc.CreateSubnetRequest) (*rpc.CreateSubnetResponse, error) {
	return &rpc.CreateSubnetResponse{}, nil
}

// Retrieve details of subnets.
func (r *Registry) GetSubnets(req *rpc.GetSubnetsRequest, resp_stream rpc.Registry_GetSubnetsServer) error {
	return nil
}

// Update a subnet's metadata but not addresses.
func (r *Registry) UpdateSubnetMeta(ctx context.Context, req *rpc.UpdateSubnetMetaRequest) (*rpc.UpdateSubnetMetaResponse, error) {
	return &rpc.UpdateSubnetMetaResponse{}, nil
}

// Assign an address from within a subnet to a user's machine.
func (r *Registry) AssignSubnetAddress(ctx context.Context, req *rpc.AssignSubnetAddressRequest) (*rpc.AssignSubnetAddressResponse, error) {
	return &rpc.AssignSubnetAddressResponse{}, nil
}

// Remove an address from within a subnet from a user's machine.
func (r *Registry) RemoveSubnetAddress(ctx context.Context, req *rpc.RemoveSubnetAddressRequest) (*rpc.RemoveSubnetAddressResponse, error) {
	return &rpc.RemoveSubnetAddressResponse{}, nil
}

// Delete a subnet.
func (r *Registry) DeleteSubnet(ctx context.Context, req *rpc.DeleteSubnetRequest) (*rpc.DeleteSubnetResponse, error) {
	return &rpc.DeleteSubnetResponse{}, nil
}

func main() {
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

	// Start GRPC server
	rpcServer := RPCServer{
		Logger: logger,
		Config: config,
	}
	if err := rpcServer.Listen(); err != nil {
		logger.Fatalf("Failed to run the RPC server "+
			"on \"%s\": %s", config.RPCAddr, err)
	}
}
