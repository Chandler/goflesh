package models

type Game struct {
	Id                      int    `json:"id"`
	Name                    string `json:"name"`
	Slug                    string `json:"slug"`
	Timezone                string `json:"timezone"`
	Registration_start_time string `json:"registration_start_time"`
	Registration_end_time   string `json:"registration_end_time"`
	Running_start_time      string `json:"running_start_time"`
	Running_end_time        string `json:"running_end_time"`
}
