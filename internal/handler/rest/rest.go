package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/CRobinDev/Gemastik/internal/handler"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/middleware"
	"github.com/CRobinDev/Gemastik/pkg/response"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	Handler    *handler.Handler
	router     *gin.Engine
	middleware middleware.Interface
}

func NewRest(handler *handler.Handler, middleware middleware.Interface) *Rest {
	return &Rest{
		Handler:    handler,
		router:     gin.Default(),
		middleware: middleware,
	}
}

func (R *Rest) UserRoute() {
	v1 := R.router.Group("/api/v1")
	v1.POST("/register", R.Handler.UserHandler.RegisterUser)
	v1.POST("/login", R.Handler.UserHandler.LoginUser)
	v1.GET("/user", R.middleware.AuthenticateUser, getLoginUser)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	R.router.Run(fmt.Sprintf(":%s", port))
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.New("unauthorized").Error(),
		})
		return
	}

	response.Success(ctx, model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "success",
		Data:    user,
	})
}
