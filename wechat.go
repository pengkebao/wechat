package wechat

import (
	"net/http"
	"sync"

	"github.com/pengkebao/wechat/cache"
	"github.com/pengkebao/wechat/context"
	"github.com/pengkebao/wechat/js"
	"github.com/pengkebao/wechat/material"
	"github.com/pengkebao/wechat/menu"
	"github.com/pengkebao/wechat/oauth"
	"github.com/pengkebao/wechat/server"
	"github.com/pengkebao/wechat/template"
	"github.com/pengkebao/wechat/user"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	Cache          cache.Cache
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Context)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs() *js.Js {
	return js.NewJs(wc.Context)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Context)
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

// GetTemplate 模板消息接口
func (wc *Wechat) GetTemplate() *template.Template {
	return template.NewTemplate(wc.Context)
}
