package controllers

import (
	"github.com/robfig/revel"
)

type Application struct {
	GorpController
}

func init() {
	revel.OnAppStart(GorpInit)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

func (c Application) Index() revel.Result {
	return c.Render()
}
