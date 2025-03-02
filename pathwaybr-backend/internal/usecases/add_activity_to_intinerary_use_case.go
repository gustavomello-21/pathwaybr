package usecases

import (
	"errors"

	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/contracts"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/ports/input"
	"github.com/gustavomello-21/pathwaybr-backend/internal/usecases/repositories"
)

type AddActivityToIntineraryUseCase struct {
	ActivityyRepository  repositories.ActivityRepositories
	IntineraryRepository repositories.IntineraryRepository
}

func NewAddActivityToIntineraryUseCase(
	activityRepository repositories.ActivityRepositories,
	intineraryRepository repositories.IntineraryRepository,
) contracts.AddActivityToIntineraryUseCase {
	return &AddActivityToIntineraryUseCase{
		ActivityyRepository:  activityRepository,
		IntineraryRepository: intineraryRepository,
	}
}

func (uc *AddActivityToIntineraryUseCase) Execute(input input.AddActivityToIntineraryInput) error {
	intineraryId, err := uc.IntineraryRepository.FindById(input.IntinerarieId)
	if err != nil {
		return err
	}

	if intineraryId == nil {
		return errors.New("Intinerary Not exist")
	}

	activity := entities.Activity{
		IntineraryID: input.IntinerarieId,
		Type:         input.Type,
		Description:  input.Description,
		StartTime:    input.StartTime,
		EndTime:      input.EndTime,
	}

	err = uc.ActivityyRepository.Save(activity)
	if err != nil {
		return err
	}

	return nil
}
