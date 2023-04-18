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

func (ss *SocialService) Create(socialReqData model.SocialCreateRequest, userID string) (*model.SocialCreateResponse, error) {
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

	return &model.SocialCreateResponse{
		ID:             newSocial.ID,
		Name:           newSocial.Name,
		SocialMediaUrl: newSocial.SocialMediaUrl,
		UserID:         newSocial.UserID,
		CreatedAt:      newSocial.CreatedAt,
		UpdatedAt:      newSocial.UpdatedAt,
	}, nil
}
