package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rapicreds-backend/src/app/infra/apierror"
	"rapicreds-backend/src/app/infra/service"
	"strconv"
)

const (
	cuilPathParam = "cuil"
	cuilLength    = 11
)

type IUserRiskController interface {
	GetUserRisk(c *gin.Context)
}

type BaseUserRiskController struct {
	userRiskService service.IUserRiskService
}

func NewBaseUserRiskController(userRiskService service.IUserRiskService) *BaseUserRiskController {
	return &BaseUserRiskController{
		userRiskService: userRiskService,
	}
}

func (b *BaseUserRiskController) GetUserRisk(c *gin.Context) {
	cuil := c.Param(cuilPathParam)

	if len(cuil) != cuilLength {
		customErr := apierror.ApiError{
			Msg:    "UserRiskController - invalid cuil length",
			Status: http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, customErr)
		return
	}

	documentAsNumbers, err := strconv.ParseUint(cuil, 10, 64)
	if err != nil {
		customErr := apierror.ApiError{
			Msg:    "UserRiskController - invalid cuil format",
			Status: http.StatusBadRequest,
			Err:    err,
		}
		c.JSON(http.StatusBadRequest, customErr)
		return
	}

	userRisk, apiErr := b.userRiskService.GetUserRisk(documentAsNumbers)
	if apiErr != nil {
		customErr := apierror.ApiError{
			Msg:    "UserRiskController - error fetching user risk from userRiskService",
			Status: apiErr.Status,
			Err:    apiErr.Err,
		}
		c.JSON(http.StatusBadRequest, customErr)
		return
	}

	c.JSON(http.StatusOK, userRisk)
}
