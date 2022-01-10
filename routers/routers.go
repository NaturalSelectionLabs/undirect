package routers

import (
	"github.com/Candinya/undirect/middleware/cors"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

// Include : Register routers
func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Allow())

	for _, opt := range options {
		opt(r)
	}

	return r
}
