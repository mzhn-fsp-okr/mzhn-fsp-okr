package grpc

import (
	"context"
	"mzhn/auth/internal/domain/entity"
	"mzhn/auth/internal/transport/grpc/converters"
	"mzhn/auth/pb/authpb"

	"github.com/samber/lo"
)

func (s *Server) Find(ctx context.Context, in *authpb.FindUserRequest) (*authpb.FindUserResponse, error) {
	user, err := s.us.Find(ctx, in.Slug)
	if err != nil {
		return nil, err
	}

	return &authpb.FindUserResponse{
		User: &authpb.UserInfo{
			Id:    user.Id,
			Email: user.Email,
			Roles: lo.Map(user.Roles, func(r entity.Role, _ int) authpb.Role {
				return converters.RoleFromE(r)
			}),
			RegisteredAt: user.CreatedAt.String(),
		},
	}, nil
}
