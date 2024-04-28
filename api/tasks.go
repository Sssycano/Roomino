package api

import (
	"net/http"
	"roomino/service"
	"roomino/types"

	"github.com/gin-gonic/gin"
)

func UnitInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UnitInforeq // 读取请求数据
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err)) // 错误处理
			return
		}

		// 创建 TaskSrv 服务实例
		taskSrv := service.GetTaskSrv()

		// 调用服务层获取数据
		resp, err := taskSrv.GetAvailableUnitsWithPetPolicy(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err)) // 错误处理
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
