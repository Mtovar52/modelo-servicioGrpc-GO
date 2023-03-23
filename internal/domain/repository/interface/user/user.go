package irepository_user

import (
	"GRPC-AUTH/internal/domain/entity"
	objectvalue "GRPC-AUTH/internal/domain/object-value/user"
)

type RepositoryUserCrud interface {
	CreateUser(user entity.User) *objectvalue.ResponseUsers
	DeleteUser(ID int64) *objectvalue.ResponseUsers
	UpdateUser(ID int64, user entity.User) *objectvalue.ResponseUsers
	ListUser(offset int32) *objectvalue.ResponseUsers
	GetByIdUser(ID int64) *objectvalue.ResponseLogin
	FindUserByEmailAndNick(email string) *objectvalue.ResponseFind
}
