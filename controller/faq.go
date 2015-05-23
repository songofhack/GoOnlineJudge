package controller

import (
	"GoOnlineJudge/class"
	"restweb"
)

type FAQController struct {
	class.Controller
} //@Controller

//faq 页面
//@URL: /api/faq @method: GET
func (fc *FAQController) Index() {
	restweb.Logger.Debug("FAQ Page")

	fc.RenderTemplate("view/layout.tpl", "view/faq.tpl")
}
