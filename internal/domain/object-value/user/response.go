package user

import pb "GRPC-AUTH/internal/infra/proto/user"

type Response struct {
	Title   string
	Message string
	IsOk    bool
	Status  int32
	Value   *pb.User
}
