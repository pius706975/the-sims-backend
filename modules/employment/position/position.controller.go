package position

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
)

type employeePositionController struct {
	service interfaces.PositionService
}

func NewEmployeePositionController(service interfaces.PositionService) *employeePositionController {
	return &employeePositionController{service}
}

// =====================================================
// Employee
// =====================================================

// CreatePosition godoc
// @Summary Create position
// @Description Create a new position with manual ID and name of the type
// @tags Employee Position
// @Accept json
// @Produce json
// @Param positionData body models.CreatePositionRequest true "Position data"
// @Param Authorization header string true "Authorization token"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /api/position/create-position [post]
func (controller *employeePositionController) CreatePosition(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	name, exists := ctx.Get("name")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Unauthorized",
		})

		return
	}

	var req models.CreatePositionRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		errors := utils.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Validation error",
			"errors":  errors,
		})

		return
	}

	positionData := models.Position{
		ID:           req.ID,
		PositionName: req.PositionName,
	}

	responseData, status := controller.service.CreatePosition(&positionData, name.(string))
	ctx.JSON(status, responseData)
}

// DeletePosition godoc
// @Summary Delete a position
// @Description Delete a position by id
// @Tags Employee Position
// @Accept json
// @Produce json
// @Param position_id path string true "Position ID"
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/position/delete/{position_id} [delete]
func (controller *employeePositionController) DeletePosition(ctx *gin.Context) {
	id := ctx.Param("position_id")

	responseData, status := controller.service.DeletePosition(id)

	ctx.JSON(status, responseData)
}

// GetPositions godoc
// @Summary Get all positions
// @Description Retrieve all position
// @Tags Employee Position
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/position/positions [get]
func (controller *employeePositionController) GetPositions(ctx *gin.Context) {
	responseData, status := controller.service.GetPositions()
	ctx.JSON(status, responseData)
}
