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
	userDao    dao.UserDao
}

func (this *AuthMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := service.GetTokenFromBearerAuth(c)
		if err {
			c.AbortWithStatusJSON(http.StatusBadRequest, data.Message{
				ApiStatus: utils.API_INCORRECT_AUTH_HEADER,
				Message:   utils.ERROR_INCORRECT_AUTH_HEADER,
			})
		} else {
			valid, userEmail := this.jwtService.VerifyToken(token)
			user := this.userDao.GetUserByEmail(userEmail)
			if valid {
				c.Set("userId", user.ID)
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, data.Message{
					ApiStatus: utils.API_INVALID_TOKEN,
					Message:   utils.ERROR_INVALID_TOKEN,
				})
			}
		}
	}
}
