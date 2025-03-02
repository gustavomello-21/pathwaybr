package v1

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/trip/dto"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
)

type TripController struct {
	createTripUseCase   contracts.CreateTripUseCase
	getUserTripsUseCase contracts.GetUserTripsUseCase
	getTripByIdUseCase  contracts.GetTripByIdUseCase
}

func NewTripController(
	createTripUseCase contracts.CreateTripUseCase,
	getUserTripsUseCase contracts.GetUserTripsUseCase,
	getTripByIdUseCase contracts.GetTripByIdUseCase,
) *TripController {
	return &TripController{
		createTripUseCase:   createTripUseCase,
		getUserTripsUseCase: getUserTripsUseCase,
		getTripByIdUseCase:  getTripByIdUseCase,
	}
}

func (t *TripController) Index(httpContext *gin.Context) {
	userId := httpContext.Param("user_id")

	parsedUserId, err := strconv.Atoi(userId)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	input := input.GetUserTripsInput{
		UserId: parsedUserId,
	}

	trips, err := t.getUserTripsUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(500, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(200, gin.H{"response": trips})
}

func (t *TripController) Show(httpContext *gin.Context) {
	tripId := httpContext.Param("trip_id")

	parsedTripId, err := strconv.Atoi(tripId)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid trip ID"})
		return
	}
	input := input.GetTripByIdInput{
		TripId: parsedTripId,
	}

	trip, err := t.getTripByIdUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(500, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(200, gin.H{"response": trip})
}

func (t *TripController) Create(httpContext *gin.Context) {
	var createTripDto dto.CreateTripDto
	if err := httpContext.ShouldBindJSON(&createTripDto); err != nil {
		httpContext.JSON(400, gin.H{"error": err})
		return
	}
	startDate, err := time.Parse("2006-01-02", createTripDto.StartDate)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD."})
		return
	}

	endDate, err := time.Parse("2006-01-02", createTripDto.EndDate)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD."})
		return
	}

	input := input.CreateTripInput{
		UserId:    createTripDto.UserId,
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = t.createTripUseCase.Execute(input)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err})
	}

	httpContext.JSON(200, gin.H{"message": "TripCreated with success"})
}
