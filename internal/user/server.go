package user

import (
	"context"

	"banka-raf/gen/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	user.UnimplementedUserServiceServer
}

func (s *Server) GetEmployeeById(ctx context.Context, req *user.GetEmployeeByIdRequest) (*user.EmployeeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
