package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"junjun-box-api/config"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ToolHandler struct {
}

func (h ToolHandler) InitRouter(g *gin.RouterGroup) {
	g.POST("upload", h.upload)
}

func (h ToolHandler) upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		JSON(c, err)
		return
	}

	u, _ := url.Parse(config.Conf.Cos.RawUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Timeout: 60 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Conf.Cos.SecretID,
			SecretKey: config.Conf.Cos.SecretKey,
		},
	})

	uid := uuid.New()
	parts := strings.Split(file.Filename, ".")
	substr := parts[len(parts)-1]
	name := "img/" + uid.String() + "." + substr

	src, err := file.Open()
	if err != nil {
		JSON(c, err)
		return
	}
	defer src.Close()

	_, err = client.Object.Put(context.Background(), name, src, nil)
	if err != nil {
		slog.Error(err.Error())
		JSON(c, err)
		return
	}

	JSON(c, name)
}
