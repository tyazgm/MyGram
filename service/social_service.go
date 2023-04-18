package service

import "MyGram/repository"

type SocialService struct {
	socialRepository repository.SocialRepository
}

func NewSocialService(socialRepository repository.SocialRepository) *SocialService {
	return &SocialService{
		socialRepository: socialRepository,
	}
}
