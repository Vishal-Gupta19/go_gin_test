package controller

import (
	"github.com/Vishal-Gupta19/go_gin_test/entity"
	"github.com/Vishal-Gupta19/go_gin_test/service"
	"github.com/gin-gonic/gin"
)

// VideoController interface
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

// New VideoController constructor
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
