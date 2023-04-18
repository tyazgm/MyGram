package controller

import "MyGram/service"

type SocialController struct {
	socialService service.SocialService
}

func NewSocialController(socialService service.SocialService) *SocialController {
	return &SocialController{
		socialService: socialService,
	}
}
