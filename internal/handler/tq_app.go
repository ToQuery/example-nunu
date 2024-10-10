package handler

import (
	apiV1 "example-nunu/api/v1"
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
	apiV1.HandleSuccess(ctx, map[string]interface{}{
		"0x00": "Hello World!",
		":)":   "Thank you for using nunu!",
	})
}
