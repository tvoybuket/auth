package api

import (
	"context"

	desc "github.com/tvoybuket/auth/pkg/auth_v1"
)

func (s *server) LoginEmail(ctx context.Context, req *desc.LoginEmailRequest) (*desc.LoginResponse, error) {
	s.logger.Info("Login called", "email", req.Email)

	return &desc.LoginResponse{AccessToken: "fsdfsdf", RefreshToken: "Login successful"}, nil
}
