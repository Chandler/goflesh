package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Game struct {
	Id                      int        `json:"id"`
	Name                    string     `json:"name"`
	Slug                    string     `json:"slug"`
	Organization_id         int        `json:"organization_id"`
	Timezone                string     `json:"timezone"`
	Registration_start_time *time.Time `json:"registration_start_time"`
	Registration_end_time   *time.Time `json:"registration_end_time"`
	Running_start_time      *time.Time `json:"running_start_time"`
	Running_end_time        *time.Time `json:"running_end_time"`
	Created                 *time.Time `json:"created"`
	Updated                 *time.Time `json:"updated"`
}

func (model *Game) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Created = &now
	model.Updated = model.Created
	return nil
}

func (model *Game) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Updated = &now
	return nil
}
