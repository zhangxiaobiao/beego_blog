package controllers

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) Index()  {
	this.TplName = "index/index.html"
}
