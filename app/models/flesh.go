package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type TimeTrackedModel struct {
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

func (model *TimeTrackedModel) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Created = &now
	model.Updated = model.Created
	return nil
}

func (model *TimeTrackedModel) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Updated = &now
	return nil
}
