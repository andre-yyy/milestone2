package topups

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
)

type TopupService struct {
	topupRepository repository.TopupRepository
}

func NewTopupService(topupRepository repository.TopupRepository) *TopupService {
	return &TopupService{topupRepository: topupRepository}
}

func (s *TopupService) GetAllTopups() ([]entity.Topup, error) {
	allTopups, err := s.topupRepository.Find()
	if err != nil {
		return []entity.Topup{}, err
	}
	return allTopups, nil
}

func (s *TopupService) GetAllTopupsByUserId(userId string) ([]entity.Topup, error) {
	allTopups, err := s.topupRepository.FindByUserId(userId)
	if err != nil {
		return []entity.Topup{}, err
	}
	return allTopups, nil
}

func (s *TopupService) GetAllTopupsByStatus(status string, userId *string) ([]entity.Topup, error) {
	allTopups, err := s.topupRepository.FindByStatus(status, userId)
	if err != nil {
		return []entity.Topup{}, err
	}
	return allTopups, nil
}

func (s *TopupService) GetTopupByExternalId(externalId string) (entity.Topup, error) {
	topup, err := s.topupRepository.FindOneByExternalId(externalId)
	if err != nil {
		return entity.Topup{}, err
	}
	return topup, nil
}

func (s *TopupService) CreateTopup(newTopup entity.Topup) (entity.Topup, error) {
	createdTopup, err := s.topupRepository.Create(newTopup)
	if err != nil {
		return entity.Topup{}, err
	}
	return createdTopup, nil
}

func (s *TopupService) UpdateTopupStatusByExternalId(topupToUpdate entity.Topup) (entity.Topup, error) {
	updatedTopup, err := s.topupRepository.UpdateStatusByExternalId(topupToUpdate)
	if err != nil {
		return entity.Topup{}, err
	}
	return updatedTopup, nil
}
