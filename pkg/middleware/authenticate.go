package middleware

import (
	"errors"
	"net/http"
	"strings"
	"github.com/CRobinDev/Gemastik/pkg/helper"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.New("unauthorized").Error(),
		})
		ctx.Abort()
	}

	token := strings.Split(bearer, " ")[1]
	userID, err := helper.ValidateToken(token)
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.New("unauthorized").Error(),
		})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", userID)
	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userID,
	})
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.New("unauthorized").Error(),
		})
		ctx.Abort()
	}

	ctx.Set("user", user)

	ctx.Next()
}
