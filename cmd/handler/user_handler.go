package handler

import (
	"GRPC-AUTH/internal/domain/entity"
	repository_user "GRPC-AUTH/internal/domain/repository/implement/user"
	irepository_user "GRPC-AUTH/internal/domain/repository/interface/user"
	"context"

	pb "GRPC-AUTH/internal/infra/proto/user"

	"gorm.io/gorm"
)

func NewServerUser(userCrud irepository_user.RepositoryUserCrud, DB *gorm.DB) *serveruser {
	return &serveruser{
		userCrud: repository_user.UserRepository(DB),
	}
}

type serveruser struct {
	userCrud irepository_user.RepositoryUserCrud
	pb.UnimplementedUserServiceServer
}

func (s *serveruser) CreateUser(context context.Context, user *pb.User) (*pb.UserResponse, error) {

	userObject := entity.User{
		Name:        user.GetName(),
		NumDocument: user.GetNumDocument(),
		Email:       user.GetEmail(),
		Phone:       user.GetPhoneContact(),
		NickName:    user.GetNickName(),
		Conditions:  user.GetConditions(),
		Password:    user.GetPassword(),
		State:       1,
	}

	response := s.userCrud.CreateUser(userObject)

	responsePB := pb.UserResponse{
		Title:   response.Title,
		IsOk:    response.IsOk,
		Message: response.Message,
		Status:  response.Status,
	}
	return &responsePB, nil

}

func (s *serveruser) UpdateUser(context context.Context, requestUser *pb.UpdateRequestUser) (*pb.UserResponse, error) {

	userObject := entity.User{
		Name:        requestUser.User.GetName(),
		NumDocument: requestUser.User.GetNumDocument(),
		NickName:    requestUser.User.GetNickName(),
	}

	response := s.userCrud.UpdateUser(requestUser.GetId(), userObject)

	responsePB := pb.UserResponse{
		Title:   response.Title,
		IsOk:    response.IsOk,
		Message: response.Message,
		Status:  response.Status,
	}

	return &responsePB, nil
}

func (s *serveruser) ListUser(context context.Context, list *pb.ListRequestUser) (*pb.ListAllResponse, error) {
	response := s.userCrud.ListUser(list.GetOffset())

	responsePB := pb.ListAllResponse{
		ListResponse: response.Value,
	}

	return &responsePB, nil
}

func (s *serveruser) DeleteUser(context context.Context, deleteRequest *pb.DeleteRequestUser) (*pb.UserResponse, error) {
	response := s.userCrud.DeleteUser(deleteRequest.GetId())

	responsePB := pb.UserResponse{
		Title:   response.Title,
		IsOk:    response.IsOk,
		Message: response.Message,
		Status:  response.Status,
	}

	return &responsePB, nil
}

func (s *serveruser) GetByIdUser(context context.Context, requestUser *pb.GetById) (*pb.UserLogin, error) {

	response := s.userCrud.GetByIdUser(requestUser.GetId())
	userResponse := response.Value
	if userResponse == nil {
		userResponse = &pb.UserLogin{}
	}
	return userResponse, nil
}

func (s *serveruser) FindUserByEmailAndNick(context context.Context, veryf *pb.FindVerifRequest) (*pb.UserLogin, error) {

	response := s.userCrud.FindUserByEmailAndNick(veryf.GetEmail())
	userResponse := response.Value
	if userResponse == nil {
		userResponse = &pb.UserLogin{}
	}
	return userResponse, nil
}
