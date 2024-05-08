package response

import (
	"github.com/CRobinDev/Gemastik/model"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data model.ServiceResponse) {
	res := &model.Response{
		Error:   data.Error,
		Message: data.Message,
		Data:    data.Data,
	}
	ctx.JSON(data.Code, res)
}

func Error(ctx *gin.Context, data model.ServiceResponse) {
	res := &model.Response{
		Error:   data.Error,
		Message: data.Message,
		Data:    data.Data,
	}

	ctx.AbortWithStatusJSON(data.Code, res)
}