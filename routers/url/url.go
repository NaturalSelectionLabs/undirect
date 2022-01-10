package url

import (
	"github.com/Candinya/undirect/handler/url"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {

	e.GET("/", url.GetHandler)

}
