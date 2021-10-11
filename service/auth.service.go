package service

import (
	"net/http"
	"savi8sant8s/api/dao"
	"savi8sant8s/api/data"
	"savi8sant8s/api/entity"
	"savi8sant8s/api/utils"
	"savi8sant8s/api/validation"
)

type AuthService struct {
	sessionDao dao.SessionDao
	userDao    dao.UserDao
}

func (this *AuthService) ValidToken(token string) (bool, uint) {
	valid := this.sessionDao.Valid(token)
	if !valid {
		return false, 0
	}
	return true, this.sessionDao.UserID(token)
}

func (this *AuthService) Register(user entity.User) (int, interface{}) {
	valid, messageError := validation.ValidUser(user)
	if !valid {
		return http.StatusBadRequest, data.Message {
			ApiStatus: utils.API_REGISTER_INCORRECT_FIELDS, 
			Message: messageError,
		}
	}
	alreadyExists := this.userDao.Exists(user.Email)
	if alreadyExists {
		return http.StatusConflict, data.Message {
			ApiStatus: utils.API_REGISTER_EMAIL_ALREADY_EXISTS, 
			Message: utils.EMAIL_ALREADY_REGISTERED,
		}
	}
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	this.userDao.New(user)
	return http.StatusOK, data.Message {
		ApiStatus: utils.API_REGISTER_SUCCESS, 
		Message: utils.USER_REGISTERED,
	}
}

func (this *AuthService) CreateSession(user data.Login) (int, interface{}) {
	valid, messageError := validation.ValidLogin(user)
	if !valid {
		return http.StatusBadRequest, data.Message {
			ApiStatus: utils.API_REGISTER_INCORRECT_FIELDS, 
			Message: messageError,
		}
	}
	exists := this.userDao.Exists(user.Email)
	if !exists {
		return http.StatusConflict, data.Message{
			ApiStatus: utils.API_LOGIN_EMAIL_NOT_REGISTERED, 
			Message: utils.EMAIL_NOT_REGISTERED,
		}
	}
	hash := this.userDao.GetHashByEmail(user.Email)
	validPassword := CheckPasswordHash(user.Password, hash)
	if !validPassword {
		return http.StatusUnauthorized, data.Message {
			ApiStatus: utils.API_LOGIN_INVALID_CREDENTIALS, 
			Message: utils.INVALID_CREDENTIALS,
		}
	}
	id := this.userDao.GetIdByEmail(user.Email)
	token := GenerateToken(30)
	this.sessionDao.CloseLastSessions(id)
	this.sessionDao.Create(entity.Session{Expired: false, Token: token, UserID: id})
	return http.StatusOK, data.MessageToken {
		Message: data.Message {
			ApiStatus: utils.API_LOGIN_SUCCESS, 
			Message: utils.LOGIN_SUCCESS,
		},
		Token: token,
	}
}

func (this *AuthService) CloseSession(userToken string) (int, interface{}) {
	this.sessionDao.Close(userToken)
	return http.StatusOK, data.Message {
		ApiStatus: utils.API_LOGOUT_SUCCESS, 
		Message: utils.LOGOUT_SUCCESS,
	}
}
