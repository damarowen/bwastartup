package middleware

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)


func AuthorizeJWT(jwtService auth.IJwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res := helper.ApiResponse(false, "failed to procces request", http.StatusBadRequest, "","no token found")
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			res := helper.ApiResponse(false, "token not valid", http.StatusBadRequest, "",err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("CurrentUser", claims)
			c.Next()
		}
	}
}

//limt upload
func BodySizeMiddleware(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, 1 * 1024 * 1024) // 1 Mb)

	c.Next()
}
