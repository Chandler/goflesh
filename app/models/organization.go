package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Organization struct {
	Id               int        `json:"id"`
	Name             string     `json:"name"`
	Slug             string     `json:"slug"`
	Default_timezone string     `json:"default_timezone"`
	Created          *time.Time `json:"created"`
	Updated          *time.Time `json:"updated"`
}

func (model *Organization) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Created = &now
	model.Updated = model.Created
	return nil
}

func (model *Organization) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Updated = &now
	return nil
}
