package repository_user

import (
	"GRPC-AUTH/internal/domain/entity"
	objectvalue "GRPC-AUTH/internal/domain/object-value/user"
	irepository_user "GRPC-AUTH/internal/domain/repository/interface/user"

	//hash "GRPC-AUTH/internal/domain/validator"

	pb "GRPC-AUTH/internal/infra/proto/user"

	validateErr "GRPC-AUTH/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type userCrud struct {
	DB *gorm.DB
}

func UserRepository(DB *gorm.DB) irepository_user.RepositoryUserCrud {
	return &userCrud{
		DB,
	}
}

// //////////////////////////    INSERT   ////////////////////////////////////
func (u *userCrud) CreateUser(user entity.User) *objectvalue.ResponseUsers {

	result := u.DB.Save(&user)

	if err := result.Error; err != nil {
		message := validateErr.ValidateError(err)
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no existoso",
			IsOk:    false,
			Message: message + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return &objectvalue.ResponseUsers{
		Title:   "¡Creado exitosamente!",
		IsOk:    true,
		Message: "El usuario se ha creado",
		Status:  http.StatusCreated,
	}
}

// //////////////////////////    DELETE   ////////////////////////////////////
func (u *userCrud) DeleteUser(ID int64) *objectvalue.ResponseUsers {

	result := u.DB.Model(&entity.User{}).Where("id", ID).Update("state", 0)

	if err := result.Error; err != nil {
		message := validateErr.ValidateError(err)
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no existoso",
			IsOk:    false,
			Message: message + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}
	if result.RowsAffected == 0 {
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no existoso",
			IsOk:    false,
			Message: "Lo sentimos Rol no eliminado ID no encontrado ",
			Status:  http.StatusBadRequest,
		}
	}
	return &objectvalue.ResponseUsers{
		Title:   "Eliminado exitosamente!",
		IsOk:    true,
		Message: "Se eliminó correctamente",
		Status:  http.StatusOK,
	}
}

// ////////////////////////////// UPDATE //////////////////////////////////////////////
func (u *userCrud) UpdateUser(ID int64, user entity.User) *objectvalue.ResponseUsers {

	whereParameter := map[string]interface{}{"id": ID}
	result := u.DB.Where(whereParameter).Updates(user)

	if err := result.Statement.Error; err != nil {
		message := validateErr.ValidateError(err)
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no existoso",
			IsOk:    false,
			Message: message + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}
	if err := result.Error; err != nil {
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no exitoso",
			IsOk:    false,
			Message: "El usuario no actualizado  " + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}
	if result.RowsAffected == 0 {
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no exitoso",
			IsOk:    false,
			Message: "El ID usado probablemente no se encuentre en el rango ",
			Status:  http.StatusBadRequest,
		}
	}
	return &objectvalue.ResponseUsers{
		Title:   "Editado exitosamente!",
		IsOk:    true,
		Message: "El usuario fue editado correctamente",
		Status:  http.StatusOK,
	}
}
func (u *userCrud) ListUser(offset int32) *objectvalue.ResponseUsers {
	var users []*pb.UserLogin

	err := u.DB.Model(&entity.User{}).Limit(int(offset)).Find(&users).Error
	if err != nil {
		return &objectvalue.ResponseUsers{
			Title:   "Proceso no existoso",
			IsOk:    false,
			Message: "No se han podido listar los usuarios" + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return &objectvalue.ResponseUsers{
		Title:   "Borrado exitosamente!",
		IsOk:    true,
		Message: "Se ha cambiado el estado de la empresa",
		Status:  http.StatusOK,
		Value:   users,
	}
}

func (u *userCrud) GetByIdUser(ID int64) *objectvalue.ResponseLogin {
	var user *pb.UserLogin
	//var user *pb.User
	whereParameter := map[string]interface{}{"id": ID, "state": 1}
	result := u.DB.Model(entity.User{}).Where(whereParameter).First(&user)

	if err := result.Error; err != nil {
		return &objectvalue.ResponseLogin{
			Title:   "Proceso no exitoso",
			IsOk:    false,
			Message: "El Id utilizado no existe " + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return &objectvalue.ResponseLogin{
		Title:   "Proceso exitosamente!",
		IsOk:    true,
		Message: "Se ha encontrado el usuario con el ID",
		Status:  http.StatusCreated,
		Value:   user,
	}
}

func (u *userCrud) FindUserByEmailAndNick(email string) *objectvalue.ResponseFind {
	var user *pb.UserLogin
	whereParameter := map[string]interface{}{"email": email, "state": 1}

	result := u.DB.Model(entity.User{}).Where(whereParameter).First(&user)

	if result.RowsAffected == 0 {
		return &objectvalue.ResponseFind{
			Title:   "Proceso no exitoso",
			IsOk:    false,
			Message: "El ID o Id_Pivot usado probablemente no se encuentre en el rango ",
			Status:  http.StatusBadRequest,
		}
	}
	if err := result.Error; err != nil {
		return &objectvalue.ResponseFind{
			Title:   "Proceso no exitoso",
			IsOk:    false,
			Message: "El Id user utilizado no existe " + err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	return &objectvalue.ResponseFind{
		Title:   "Proceso exitosamente!",
		IsOk:    true,
		Message: "Se ha encontrado el usuario con el ID",
		Status:  http.StatusCreated,
		Value:   user,
	}
}
