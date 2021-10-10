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
	validation validation.AuthValidation
	sessionDao dao.SessionDao
	userDao dao.UserDao
}

func (service *AuthService) ValidToken(token string) (bool, uint) {
	valid := service.sessionDao.Valid(token)
	if !valid {
		return false, 0
	}
	return true, service.sessionDao.UserID(token)
}

func (service *AuthService) RegisterUser(user entity.User) (int, data.Message) {
	valid, messageError := service.validation.ValidRegister(user)
	if !valid {
		return MakeRes(http.StatusBadRequest, API_REGISTER_INCORRECT_FIELDS, messageError)
	} 
	alreadyExists := service.userDao.Exists(user.Email)
	if alreadyExists {
		return MakeRes(http.StatusConflict, API_REGISTER_EMAIL_ALREADY_EXISTS, utils.MESSAGE_EMAIL_ALREADY_REGISTERED)
	}
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	service.userDao.New(user)
	return MakeRes(http.StatusOK, API_REGISTER_SUCCESS, utils.MESSAGE_USER_REGISTERED)
}

func (service *AuthService) Login(user data.Login) (int, data.Message) {
	valid, messageError := service.validation.ValidLogin(user)
	if !valid {
		return MakeRes(http.StatusBadRequest,API_REGISTER_INCORRECT_FIELDS, messageError)
	} 
	exists := service.userDao.Exists(user.Email)
	if !exists {
		return MakeRes(http.StatusConflict, API_LOGIN_EMAIL_NOT_REGISTERED, utils.MESSAGE_EMAIL_NOT_REGISTERED)
	}
	hash := service.userDao.GetHashByEmail(user.Email)
	validPassword := CheckPasswordHash(user.Password, hash)
	if !validPassword {
		return MakeRes(http.StatusUnauthorized, API_LOGIN_INVALID_CREDENTIALS, utils.MESSAGE_INVALID_CREDENTIALS)
	}
	id := service.userDao.GetIdByEmail(user.Email)
	token := GenerateToken(30)
	service.sessionDao.Create(entity.Session{Expired: false, Token: token, UserID: id})
	return MakeRes(http.StatusOK, API_LOGIN_SUCCESS, utils.MESSAGE_LOGIN_SUCCESS)
}

func (this *AuthService) Logout(userToken string) (int, data.Message) {
	this.sessionDao.Close(userToken)	
	return MakeRes(http.StatusOK, API_LOGOUT_SUCCESS, utils.MESSAGE_LOGOUT_SUCCESS)
}