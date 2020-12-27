package controllers

import (
	"github.com/afocus/captcha"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"image/color"
	"image/png"
)

type CaptchaController struct {
	beego.Controller
}

var cap *captcha.Captcha

// @router /captcha [get]
func (this *CaptchaController) Get()  {
	id := this.GetString("t")
	cap = captcha.New()
	if err := cap.SetFont("static/font/comic.ttf"); err != nil {
		panic(err.Error())
	}
	cap.SetSize(110, 38)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	img, str := cap.Create(6, captcha.CLEAR)
	logs.Info(str)
	this.SetSession(id,str)
	png.Encode(this.Ctx.ResponseWriter, img)
}
