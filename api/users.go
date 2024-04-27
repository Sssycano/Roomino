package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roomino/ctl"
	"roomino/e"
	"roomino/service"
	"roomino/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserSrv()
			resp, err := l.Register(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}
func ErrorResponse(err error) *ctl.TrackedErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range ve {
			field := fieldError.Field()
			tag := fieldError.Tag()
			return ctl.RespError(err, fmt.Sprintf("%s validation failed: %s", field, tag))
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, "JSON type mismatch", e.InvalidParams)
	}
	return ctl.RespError(err, "Invalid parameters", e.InvalidParams)
}
