package api

import (
	desc "github.com/tvoybuket/auth/pkg/auth_v1"
	"github.com/tvoybuket/tblib/tblogger"
)

type server struct {
	desc.UnimplementedAuthServiceServer
	logger *tblogger.Logger
}

func NewServer(logger *tblogger.Logger) *server {
	return &server{
		logger: logger,
	}
}
