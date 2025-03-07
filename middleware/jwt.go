package middleware

import (
	"doovvvblog/utils"
	"doovvvblog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey []byte = []byte(utils.AppConfig.Server.Secret)

func CreateJwt(username string, password string) (string, error) {
	cliaims := jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Now().Add(10 * time.Hour).Unix(),
		"iss":      "doovvv",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliaims)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
func verifyJwt(tokenString string) (jwt.MapClaims, int) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, errmsg.ERROR
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, errmsg.SUCCESS
	}
	return nil, errmsg.ERROR
}
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// code := errmsg.SUCCESS
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    errmsg.ERROR_TOKEN_NOT_EXIST,
				"message": errmsg.GetErrorMsg(errmsg.ERROR_TOKEN_NOT_EXIST),
			})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    errmsg.ERROR_TOKEN_TYPE_WRONG,
				"message": errmsg.GetErrorMsg(errmsg.ERROR_TOKEN_TYPE_WRONG),
			})
			c.Abort()
			return
		}
		claims, code := verifyJwt(tokenString)
		if code == errmsg.ERROR {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		expFloat, _ := claims["exp"].(float64)
		if time.Now().Unix() > int64(expFloat) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    errmsg.ERROR_TOKEN_RUNTIME,
				"message": errmsg.GetErrorMsg(errmsg.ERROR_TOKEN_RUNTIME),
			})
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Next()
	}
}
