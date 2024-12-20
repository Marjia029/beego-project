package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// BaseController is the common controller for shared logic
type BaseController struct {
	web.Controller
}

func (c *BaseController) IsAjax() bool {
	return c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest"
}
