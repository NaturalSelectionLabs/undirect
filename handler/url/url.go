package url

import (
	"github.com/Candinya/undirect/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func GetHandler(ctx *gin.Context) {
	encodedUrl := ctx.Query("url")
	decodedUrl, err := url.QueryUnescape(encodedUrl)
	if err != nil {
		handler.HandleError(ctx, err.Error(), http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	_, err = url.ParseRequestURI(decodedUrl)
	if err != nil {
		handler.HandleError(ctx, err.Error(), http.StatusBadRequest)
		ctx.Abort()
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", decodedUrl, nil)
	if err != nil {
		handler.HandleError(ctx, err.Error(), http.StatusInternalServerError)
		ctx.Abort()
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		handler.HandleError(ctx, err.Error(), http.StatusInternalServerError)
		ctx.Abort()
		return
	}
	finalUrl := resp.Request.URL.String()
	handler.HandleSuccess(ctx, finalUrl, http.StatusOK)
}
