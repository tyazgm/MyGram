package service

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/repository"
	"time"
)

type SocialService struct {
	socialRepository repository.SocialRepository
}

func NewSocialService(socialRepository repository.SocialRepository) *SocialService {
	return &SocialService{
		socialRepository: socialRepository,
	}
}

func (ss *SocialService) Create(socialReqData model.SocialCreateRequest, userID string) (*model.SocialResponse, error) {
	socialID := helper.GenerateID()
	newSocial := model.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaUrl: socialReqData.SocialMediaUrl,
		UserID:         userID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := ss.socialRepository.Create(newSocial)
	if err != nil {
		return nil, err
	}

	return &model.SocialResponse{
		ID:             newSocial.ID,
		Name:           newSocial.Name,
		SocialMediaUrl: newSocial.SocialMediaUrl,
		UserID:         newSocial.UserID,
		CreatedAt:      newSocial.CreatedAt,
		UpdatedAt:      newSocial.UpdatedAt,
	}, nil
}

func (ss *SocialService) GetAll() ([]model.SocialResponse, error) {
	photosResult, err := ss.socialRepository.FindAll()
	if err != nil {
		return []model.SocialResponse{}, err
	}

	socialResponse := []model.SocialResponse{}
	for _, socialMedia := range photosResult {
		socialResponse = append(socialResponse, model.SocialResponse(socialMedia))
	}

	return socialResponse, nil
}

func (ss *SocialService) GetOne(socialID string) (model.SocialResponse, error) {
	socialMediaResponse, err := ss.socialRepository.FindByID(socialID)
	if err != nil {
		return model.SocialResponse{}, err
	}

	return model.SocialResponse(socialMediaResponse), nil
}
