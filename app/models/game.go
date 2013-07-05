package models

import (
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
	TimeTrackedModel
}
