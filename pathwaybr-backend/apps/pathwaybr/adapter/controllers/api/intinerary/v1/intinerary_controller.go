package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/intinerary/dto"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type IntineraryController struct {
	addIntineraryToTripUseCase contracts.AddIntineraryToTripUseCase
}

func NewIntineraryController(addIntineraryToTripUseCase contracts.AddIntineraryToTripUseCase) *IntineraryController {
	return &IntineraryController{
		addIntineraryToTripUseCase: addIntineraryToTripUseCase,
	}
}

func (i *IntineraryController) Create(httpContext *gin.Context) {
	tripID := httpContext.Param("trip_id")
	parsedTripId, err := strconv.Atoi(tripID)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid trip ID"})
		return
	}
	var createIntineraryDto dto.CreateIntineraryDto
	if err := httpContext.ShouldBindJSON(&createIntineraryDto); err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input := input.AddIntineraryToTripInput{
		TripId:    parsedTripId,
		DayNumber: createIntineraryDto.DayNumber,
	}

	err = i.addIntineraryToTripUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err})
	}

	httpContext.JSON(200, gin.H{"response": "intinerary Created"})
}
