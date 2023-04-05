package gapi

import (
	"fmt"

	db "github.com/lushenle/mundilfari/db/sqlc"
	"github.com/lushenle/mundilfari/pb"
	"github.com/lushenle/mundilfari/token"
	"github.com/lushenle/mundilfari/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedMundilfariServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC Server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
