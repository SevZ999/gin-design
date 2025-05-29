package controller

import (
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	srv *service.AccessService
}

func NewAccessController(srv *service.AccessService) *AccessController {
	return &AccessController{srv: srv}
}

func (ctrl *AccessController) GetAccess(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(http.StatusBadRequest, "invalid id"))
		return
	}
	user, err := ctrl.srv.GetAccess(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(user))
}
