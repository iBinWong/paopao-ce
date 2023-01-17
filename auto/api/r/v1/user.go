// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v3.0.1

package v1

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type User interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Logout() mir.Error
	Login(*LoginReq) (*LoginResp, mir.Error)

	mustEmbedUnimplementedUserServant()
}

type UserBinding interface {
	BindLogin(*gin.Context) (*LoginReq, mir.Error)

	mustEmbedUnimplementedUserBinding()
}

type UserRender interface {
	RenderLogout(*gin.Context, mir.Error)
	RenderLogin(*gin.Context, *LoginResp, mir.Error)

	mustEmbedUnimplementedUserRender()
}

// RegisterUserServant register User servant to gin
func RegisterUserServant(e *gin.Engine, s User, b UserBinding, r UserRender) {
	router := e.Group("r/v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderLogout(c, s.Logout())
	})

	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
			return
		}
		resp, err := s.Login(req)
		r.RenderLogin(c, resp, err)
	})

}

// UnimplementedUserServant can be embedded to have forward compatible implementations.
type UnimplementedUserServant struct {
}

func (UnimplementedUserServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedUserServant) Logout() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedUserServant) Login(req *LoginReq) (*LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedUserServant) mustEmbedUnimplementedUserServant() {}

// UnimplementedUserRender can be embedded to have forward compatible implementations.
type UnimplementedUserRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (r *UnimplementedUserRender) RenderLogout(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedUserRender) RenderLogin(c *gin.Context, data *LoginResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedUserRender) mustEmbedUnimplementedUserRender() {}

// UnimplementedUserBinding can be embedded to have forward compatible implementations.
type UnimplementedUserBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

func (b *UnimplementedUserBinding) BindLogin(c *gin.Context) (*LoginReq, mir.Error) {
	obj := new(LoginReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedUserBinding) mustEmbedUnimplementedUserBinding() {}
