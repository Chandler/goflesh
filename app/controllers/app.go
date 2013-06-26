package controllers

import (
	"github.com/robfig/revel"
)

type Application struct {
	*revel.Controller
}

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

func (c Application) Index() revel.Result {
	return c.Render()
}
