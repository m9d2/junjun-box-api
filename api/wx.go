package api

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	mConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"junjun-box-api/config"
	"junjun-box-api/model"
)

type WxHandler struct {
}

func (h WxHandler) InitRouter(g *gin.RouterGroup) {
	g.GET("wx/code2session", h.code2session)
}

func (h WxHandler) code2session(g *gin.Context) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &mConfig.Config{
		AppID:     config.Conf.Wx.AppID,
		AppSecret: config.Conf.Wx.AppSecret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetAuth()
	jsCode := g.Query("jsCode")
	res, err := a.Code2Session(jsCode)
	member := &model.Member{
		OpenId:    res.OpenID,
		AvatarUrl: "",
		NickName:  "",
	}
	if err != nil {
		JSON(g, err)
	} else {
		JSON(g, member)
	}
}
