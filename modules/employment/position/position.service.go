package position

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/the-sims-backend/interfaces"
	"github.com/pius706975/the-sims-backend/package/database/models"
	"github.com/pius706975/the-sims-backend/package/utils"
)

type employeePositionService struct {
	repo interfaces.PositionRepo
}

func NewEmployeePositionService(repo interfaces.PositionRepo) *employeePositionService {
	return &employeePositionService{repo}
}

// =====================================================
// Position
// =====================================================
func (service *employeePositionService) CreatePosition(positionData *models.Position, decodedCreatorName string) (gin.H, int) {

	positionData.CreatedAt = utils.GetCurrentTime()
	positionData.CreatedBy = decodedCreatorName

	existingPosition, err := service.repo.GetExistingPosition(positionData.ID, positionData.PositionName)

	if err != nil {
		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	if existingPosition != nil {
		return gin.H{
			"status":  400,
			"message": "Position with the ID or Name already exists",
		}, 400
	}

	newData, err := service.repo.CreatePosition(positionData)
	if err != nil {
		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	return gin.H{
		"status":  201,
		"message": "Position created successfully",
		"data":    newData,
	}, 201
}

func (service *employeePositionService) DeletePosition(id string) (gin.H, int) {
	position, err := service.repo.GetPositionById(id)

	if err != nil {
		return gin.H{
			"status":  500,
			"message": err.Error(),
		}, 500
	}

	if position == nil {
		return gin.H{
			"status":  404,
			"message": "Position not found",
		}, 404
	}

	if err := service.repo.DeletePosition(id); err != nil {
		return gin.H{
			"status":  500,
			"message": "Failed to delete position",
		}, 500
	}

	return gin.H{
		"status":  200,
		"message": "Position deleted successfully",
	}, 200

}

func (service *employeePositionService) GetPositions() (gin.H, int) {
	positions, err := service.repo.GetPositions()

	if err != nil {
		return gin.H{
			"status":  404,
			"Message": err.Error(),
		}, 404
	}

	return gin.H{
		"status":  200,
		"message": "All positions retrieved successfully",
		"data":    positions,
	}, 200
}
