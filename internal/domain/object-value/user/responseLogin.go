package user

import pb "GRPC-AUTH/internal/infra/proto/user"

type ResponseLogin struct {
	Title   string
	Message string
	IsOk    bool
	Status  int32
	Value   *pb.UserLogin
}
