package handler

import (
	"example-nunu/internal/service"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	*Handler
	appService service.AppService
}

func NewAppHandler(
	handler *Handler,
	appService service.AppService,
) *AppHandler {
	return &AppHandler{
		Handler:    handler,
		appService: appService,
	}
}

func (h *AppHandler) GetApp(ctx *gin.Context) {

}
