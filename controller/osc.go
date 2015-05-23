package controller

import (
	"GoOnlineJudge/class"
	"restweb"
)

type OSCController struct {
	class.Controller
} //@Controller

//@URL: /api/osc @method: GET
func (oc *OSCController) Index() {
	restweb.Logger.Debug("OSC Page")

	oc.RenderTemplate("view/layout.tpl", "view/osc.tpl")
}
