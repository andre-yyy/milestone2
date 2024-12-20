package activities

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
)

type ActivityService struct {
	activityRepository repository.ActivityRepository
}

func NewActivityService(activityRepository repository.ActivityRepository) *ActivityService {
	return &ActivityService{activityRepository: activityRepository}
}

func (s *ActivityService) GetAllActivities() ([]entity.Activity, error) {
	activities, err := s.activityRepository.Find()
	if err != nil {
		return []entity.Activity{}, err
	}
	return activities, nil
}

func (s *ActivityService) GetAllActivitiesByUserId(userId string) ([]entity.Activity, error) {
	activities, err := s.activityRepository.FindByUserId(userId)
	if err != nil {
		return []entity.Activity{}, err
	}
	return activities, nil
}
