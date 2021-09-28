package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/handler"
	"goapi/pkg/errno"
	"log"
	"net/http"
)

func Create(ctx *gin.Context) {
	var r UserRequest

	var err error
	if err := ctx.Bind(&r); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Printf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found "))

		handler.SendResponse(ctx,err,nil)
		return
	}


	if r.Password == "" {
		handler.SendResponse(ctx,fmt.Errorf("password is empty"),nil)
		return
	}

	handler.SendResponse(ctx,nil,UserResponse{r.Username})
	return

}
