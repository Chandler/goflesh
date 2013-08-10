package controllers

import (
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Application struct {
	GorpController
}

func init() {
	revel.OnAppStart(models.Init)
	revel.OnAppStart(GorpInit)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

func (c Application) Index() revel.Result {
	return c.Render()
}
