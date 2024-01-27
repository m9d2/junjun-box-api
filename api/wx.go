package api

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	mConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"junjun-box-api/config"
	"junjun-box-api/model"
	"junjun-box-api/service"
	"time"
)

type WxHandler struct {
	memberService service.MemberService
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
	now := time.Now()
	member := &model.Member{
		OpenId:        res.OpenID,
		LastLoginIp:   g.RemoteIP(),
		LastLoginTime: now,
		AvatarUrl:     "",
		NickName:      "",
	}
	m := h.memberService.GetMember(res.OpenID)
	if m == nil {
		member.CreateTime = now
		member.NickName = "微信用户"
		member = h.memberService.Save(member)
	} else {
		h.memberService.UpdateMember(member)
	}

	if err != nil {
		JSON(g, err)
	} else {
		JSON(g, member)
	}
}
