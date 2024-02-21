package api

import (
	"github.com/gin-gonic/gin"
	"junjun-box-api/model"
	"junjun-box-api/service"
)

type MemberHandler struct {
	service service.MemberService
}

func (h MemberHandler) InitRouter(g *gin.RouterGroup) {
	g.GET("member", h.getMember)
	g.PUT("member", h.updateMember)
}

func (h MemberHandler) getMember(g *gin.Context) {
	openid := g.Query("openid")
	m := h.service.GetMember(openid)
	JSON(g, m)
}

func (h MemberHandler) updateMember(g *gin.Context) {
	var m model.Member
	err := g.ShouldBind(&m)
	if err != nil {
		JSON(g, err)
	} else {
		h.service.UpdateMember(&m)
		JSON(g, m)
	}

}
