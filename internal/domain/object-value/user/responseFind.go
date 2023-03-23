package user

import pb "GRPC-AUTH/internal/infra/proto/user"

type ResponseFind struct {
	Title   string
	Message string
	IsOk    bool
	Status  int32
	Value   *pb.UserLogin
}
