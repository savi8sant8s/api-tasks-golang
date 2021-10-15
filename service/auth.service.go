package service

import (
	"net/http"
	"savi8sant8s/gotasks/dao"
	"savi8sant8s/gotasks/data"
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/utils"
	"savi8sant8s/gotasks/validation"
)

type AuthService struct {
	userDao    dao.UserDao
	jwtService JwtService
}

func (this *AuthService) RegisterUser(user entity.User) (int, interface{}) {
	valid, messageError := validation.ValidUser(user)
	if !valid {
		return http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_REGISTER_INCORRECT_FIELDS,
			Message:   messageError,
		}
	}
	alreadyExists := this.userDao.Exists(user.Email)
	if alreadyExists {
		return http.StatusConflict, data.Message{
			ApiStatus: utils.API_REGISTER_EMAIL_ALREADY_EXISTS,
			Message:   utils.EMAIL_ALREADY_REGISTERED,
		}
	}
	hash := EncryptPassword(user.Password)
	user.Password = hash
	this.userDao.New(user)
	return http.StatusOK, data.Message{
		ApiStatus: utils.API_REGISTER_SUCCESS,
		Message:   utils.USER_REGISTERED,
	}
}

func (this *AuthService) CreateSession(user entity.User) (int, interface{}) {
	valid, messageError := validation.ValidLogin(user)
	if !valid {
		return http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_REGISTER_INCORRECT_FIELDS,
			Message:   messageError,
		}
	}
	exists := this.userDao.Exists(user.Email)
	if !exists {
		return http.StatusConflict, data.Message{
			ApiStatus: utils.API_LOGIN_EMAIL_NOT_REGISTERED,
			Message:   utils.EMAIL_NOT_REGISTERED,
		}
	}
	userDb := this.userDao.GetUserByEmail(user.Email)
	if !ValidPassword(user.Password, userDb.Password) {
		return http.StatusUnauthorized, data.Message{
			ApiStatus: utils.API_LOGIN_INVALID_CREDENTIALS,
			Message:   utils.INVALID_CREDENTIALS,
		}
	}
	token := this.jwtService.CreateToken(user.Email)
	return http.StatusOK, data.MessageToken{
		Message: data.Message{
			ApiStatus: utils.API_LOGIN_SUCCESS,
			Message:   utils.LOGIN_SUCCESS,
		},
		Token: token,
	}
}
