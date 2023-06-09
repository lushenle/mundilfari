package gapi

import (
	"fmt"

	db "github.com/lushenle/mundilfari/db/sqlc"
	"github.com/lushenle/mundilfari/pb"
	"github.com/lushenle/mundilfari/token"
	"github.com/lushenle/mundilfari/util"
	"github.com/lushenle/mundilfari/worker"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedMundilfariServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC Server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
