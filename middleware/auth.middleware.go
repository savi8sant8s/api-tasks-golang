package middleware

import (
	"net/http"
	"savi8sant8s/gotasks/dao"
	"savi8sant8s/gotasks/data"
	"savi8sant8s/gotasks/service"
	"savi8sant8s/gotasks/utils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService service.JwtService
	authService    service.AuthService
	userDao dao.UserDao
}

func (this *AuthMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, valid := this.authService.ValiTokenHeader(c)
		if valid {
			valid, userEmail := this.jwtService.VerifyToken(token)
			user := this.userDao.FindUserByEmail(userEmail)
			if valid {
				c.Set("userId", user.ID)
				c.Next()
			} else {
				c.JSON(http.StatusBadRequest, data.Message{
					ApiStatus: utils.API_INVALID_TOKEN,
					Message:   utils.ERROR_INVALID_TOKEN,
				})
			}
		}
	}
}
