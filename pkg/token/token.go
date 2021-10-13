package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goapi/pkg/errno"
	"time"
)

type Context struct {
	ID       uint64
	Username string
}

//登录获得token
func Sign(ctx *gin.Context, context *Context) (tokenString string, err error) {
	secret := viper.GetString("jwt_secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       context.ID,
		"username": context.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))
	return
}

func ParseRequest(ctx *gin.Context) (*Context, error) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if len(tokenString) == 0 {
		return &Context{}, errno.ErrMissingTokenString
	}

	secret := viper.GetString("jwt_secret")

	return Parse(tokenString, secret)
}

func Parse(tokenString, secret string) (*Context, error) {
	context := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return context, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		context.ID = uint64(claims["id"].(float64))
		context.Username = claims["username"].(string)
		return context, nil
	} else {
		return context, err
	}
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
