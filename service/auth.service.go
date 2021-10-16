package service

import (
	"net/http"
	"savi8sant8s/gotasks/dao"
	"savi8sant8s/gotasks/data"
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/utils"
	"savi8sant8s/gotasks/validation"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	userDao    dao.UserDao
	jwtService JwtService
}

func (this *AuthService) RegisterUser(c *gin.Context, user entity.User){
	valid := this.ValidUserBody(c, user) && this.ValidUniqueUser(c, user.Email)
	if valid {
		user.Password = EncryptPassword(user.Password)
		this.userDao.Create(user)
		c.JSON(http.StatusOK, data.Message{
			ApiStatus: utils.API_REGISTER_SUCCESS,
			Message:   utils.USER_REGISTERED,
		})
	}	
}

func (this *AuthService) CreateSession(c *gin.Context, user entity.User) {
	valid := this.ValidLoginBody(c, user) && this.ValidUserEmail(c, user.Email) && this.ValidUserPassword(c, user.Email, user.Password)
	if valid {
		token := this.jwtService.CreateToken(user.Email)
		c.JSON(http.StatusOK, data.MessageToken{
			Message: data.Message{
				ApiStatus: utils.API_LOGIN_SUCCESS,
				Message:   utils.LOGIN_SUCCESS,
			},
			Token: token,
		})
	}
}

func (this *AuthService) ValidUserBody(c *gin.Context, user entity.User) bool {
	valid, messageError := validation.ValidUser(user)
	if !valid {
		c.JSON(http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_REGISTER_INCORRECT_FIELDS,
			Message:   messageError,
		})
		return false
	}
	return true
}

func (this *AuthService) ValidLoginBody(c *gin.Context, user entity.User) bool {
	valid, messageError := validation.ValidLogin(user)
	if !valid {
		c.JSON(http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_REGISTER_INCORRECT_FIELDS,
			Message:   messageError,
		})
		return false
	}
	return true
}


func (this *AuthService) ValidUniqueUser(c *gin.Context, email string) bool {
	alreadyExists := this.userDao.Exists(email)
	if alreadyExists {
		c.JSON(http.StatusConflict, data.Message{
			ApiStatus: utils.API_REGISTER_EMAIL_ALREADY_EXISTS,
			Message:   utils.EMAIL_ALREADY_REGISTERED,
		})
		return false
	}
	return true
}

func (this *AuthService) ValidUserEmail(c *gin.Context, email string) bool {
	exists := this.userDao.Exists(email)
	if !exists {
		c.JSON(http.StatusConflict, data.Message{
			ApiStatus: utils.API_LOGIN_EMAIL_NOT_REGISTERED,
			Message:   utils.EMAIL_NOT_REGISTERED,
		})
		return false
	}
	return true
}

func (this *AuthService) ValidUserPassword(c *gin.Context, email string, password string) bool {
	user := this.userDao.FindUserByEmail(email)
	if !ValidPassword(password, user.Password) {
		c.JSON(http.StatusUnauthorized, data.Message{
			ApiStatus: utils.API_LOGIN_INVALID_CREDENTIALS,
			Message:   utils.INVALID_CREDENTIALS,
		})
		return false
	}
	return true
}

func (this *AuthService) ValiTokenHeader(c *gin.Context) (string, bool) {
	token, err := GetTokenFromBearerAuth(c)
	if err {
		return "", false
	}
	c.JSON(http.StatusBadRequest, data.Message{
		ApiStatus: utils.API_INCORRECT_AUTH_HEADER,
		Message:   utils.ERROR_INCORRECT_AUTH_HEADER,
	})
	return token, true
}
