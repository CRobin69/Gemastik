package rest

import (
	"net/http"

	"github.com/CRobinDev/Gemastik/internal/handler"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/middleware"
	"github.com/CRobinDev/Gemastik/pkg/response"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	Handler    *handler.Handler
	router     *gin.Engine
	middleware middleware.IMiddleware
}

func NewRest(handler *handler.Handler, middleware middleware.IMiddleware) *Rest {
	return &Rest{
		Handler:    handler,
		router:     gin.Default(),
		middleware: middleware,
	}
}

func (r *Rest) RestRoute() {
	r.router.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusNotFound,
			Error:   true,
			Message: errors.ErrRouteNotFound.Error(),
		})
	})
	r.router.Use(middleware.CORSMiddleware())
	v1 := r.router.Group("/api/v1")
	r.LeaderboardRoute(v1)
	r.UserRoute(v1)
	r.EducationRoute(v1)

	r.router.Run()
}
func (r *Rest) EducationRoute(router *gin.RouterGroup) {
	educationRoute := router.Group("/education")
	educationRoute.POST("/chatbot", r.middleware.AuthenticateUser, r.Handler.ChatHandler.GenerateResponse)
}

func (r *Rest) LeaderboardRoute(router *gin.RouterGroup) {
	boardRoute := router.Group("board")
	boardRoute.POST("/leaderboard", r.Handler.LeaderboardHandler.CreateLeaderboard)
	boardRoute.GET("/leaderboard/all", r.Handler.LeaderboardHandler.GetLeaderboard)
	boardRoute.GET("/leaderboard/:id", r.Handler.LeaderboardHandler.GetLeaderboardByID)
}

func (r *Rest) UserRoute(router *gin.RouterGroup) {
	userRoute := router.Group("/user")
	userRoute.POST("/register", r.Handler.UserHandler.RegisterUser)
	userRoute.POST("/login", r.Handler.UserHandler.LoginUser)
	userRoute.GET("/me", r.middleware.AuthenticateUser, getLoginUser)
	userRoute.PATCH("/update-password", r.middleware.AuthenticateUser, r.Handler.UserHandler.UpdatePassword)
	userRoute.PATCH("/update-profile", r.middleware.AuthenticateUser, r.Handler.UserHandler.UpdateProfile)
	userRoute.POST("/upload-photo", r.middleware.AuthenticateUser, r.Handler.UserHandler.UploadPhoto)
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.ErrUnathorized.Error(),
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
