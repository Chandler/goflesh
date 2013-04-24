package controllers

import (
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Organizations struct {
	*revel.Controller
}

func (c Organizations) List() revel.Result {
	return GetList(models.Organization{})
}
