package handler

import (
	"example-nunu/internal/service"
	"github.com/gin-gonic/gin"
)

type TqAppHandler struct {
	*Handler
	tqAppService service.TqAppService
}

func NewTqAppHandler(
	handler *Handler,
	tqAppService service.TqAppService,
) *TqAppHandler {
	return &TqAppHandler{
		Handler:      handler,
		tqAppService: tqAppService,
	}
}

func (h *TqAppHandler) GetTqApp(ctx *gin.Context) {

}
