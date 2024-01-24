package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"junjun-box-api/config"
	"junjun-box-api/model"
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
	g.GET("weather", h.weather)
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

func (h ToolHandler) weather(c *gin.Context) {
	client := resty.New()
	weather := model.Weather{}
	city := c.Query("city")
	u := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?city=%s&key=31bf2900a2b326ff8a0e3d3d24d66bbf", city)
	resp, err := client.R().SetResult(&weather).Get(u)
	if err != nil {
		Fail(c, err)
		return
	}
	if resp.StatusCode() != 200 {
		Fail(c, errors.New(resp.Error().(string)))
		return
	}
	if weather.Status != "1" {
		Fail(c, errors.New(weather.Info))
		return
	}
	l := weather.Lives[0]
	JSON(c, l)
}
