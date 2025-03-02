package v1

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/activity/v1/dto"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type ActivityController struct {
	addActivityToIntineraryUseCase contracts.AddActivityToIntineraryUseCase
}

func NewActivityController(
	addActivityToIntineraryUseCase contracts.AddActivityToIntineraryUseCase,
) *ActivityController {
	return &ActivityController{
		addActivityToIntineraryUseCase: addActivityToIntineraryUseCase,
	}
}

func (a *ActivityController) Create(httpContext *gin.Context) {
	intineraryId := httpContext.Param("intinerary_id")
	parsedIntineraryId, err := strconv.Atoi(intineraryId)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid Intinerary ID"})
		return
	}

	var createActivityDto dto.CreateActivityDto
	err = httpContext.ShouldBindJSON(&createActivityDto)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err})
		return
	}
	startDate, err := time.Parse("2006-01-02", createActivityDto.StartTime)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD."})
		return
	}

	endDate, err := time.Parse("2006-01-02", createActivityDto.EndTime)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD."})
		return
	}

	input := input.AddActivityToIntineraryInput{
		IntinerarieId: parsedIntineraryId,
		Type:          createActivityDto.Type,
		Description:   createActivityDto.Description,
		StartTime:     startDate,
		EndTime:       endDate,
	}

	err = a.addActivityToIntineraryUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "failed to add activity"})
		return
	}

	httpContext.JSON(200, gin.H{"response": "acitivy added with success"})
}
