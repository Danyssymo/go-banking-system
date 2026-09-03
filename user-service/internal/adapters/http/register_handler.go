package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Danyssymo/go-banking-system/user-service/internal/usecase"
)

type RegisterHandler struct {
	uc *usecase.RegisterUseCase
}

func NewRegisterHandler(uc *usecase.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{uc: uc}
}

func (h *RegisterHandler) Handle(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.Execute(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		status, msg := mapError(err)
		c.JSON(status, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusCreated, newRegisterResponse(user))
}
