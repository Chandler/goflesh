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

func GameFromId(id int) (*Game, error) {
	game, err := Dbm.Get(Game{}, id)
	return game.(*Game), err
}

func (m *Game) Save() error {
	_, err := Dbm.Update(m)
	return err
}

func (g *Game) IsRunning() bool {
	now := time.Now()
	return g.Running_start_time.Before(now) && g.Running_end_time.After(now)
}

func (g *Game) TimeToStarve() time.Duration {
	// TODO: move this into a game-specific setting in the DB
	duration, _ := time.ParseDuration("72h")
	return duration
}

func (g *Game) TimeToReveal() *time.Time {
	// TODO: move this into a game-specific setting in the DB
	undercoverTime, _ := time.ParseDuration("24h")
	time := g.Running_start_time.Add(undercoverTime)
	return &time
}
