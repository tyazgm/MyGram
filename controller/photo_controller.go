package controller

import "MyGram/service"

type PhotoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: photoService,
	}
}
