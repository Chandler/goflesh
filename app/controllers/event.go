package controllers

import (
	"flesh/app/models"
	"fmt"
	"github.com/robfig/revel"
	"sort"
	"time"
)

type Events struct {
	AuthController
}

type Dated interface {
	Date() *time.Time
}

type DatedSortable []Dated

func (s DatedSortable) Len() int { return len(s) }

func (s DatedSortable) Less(i, j int) bool {
	return s[i].Date().After(*s[j].Date())
}
func (s DatedSortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ClientTagEvent struct {
	Id       string     `json:"id"`
	Type     string     `json:"type"`
	SortDate *time.Time `json:"-"`
	Tag      models.Tag `json:"tag"`
}

func (c ClientTagEvent) Date() *time.Time {
	return c.SortDate
}

type ClientPlayerEvent struct {
	Id       string     `json:"id"`
	Type     string     `json:"type"`
	SortDate *time.Time `json:"-"`
	Player   int `json:"player_id"`
}

func (c ClientPlayerEvent) Date() *time.Time {
	return c.SortDate
}

type IdWrapper struct {
	Id int
}

func (c *Events) GetTagEvents(ids_string string) DatedSortable {
	query := `
		SELECT tag.*
		FROM tag
		INNER JOIN event_tag
			ON event_tag.tag_id = tag.id
		WHERE event_tag.id IN (` + ids_string + `)
    `
	var list []*models.Tag
	_, err := Dbm.Select(&list, query)
	if err != nil {
		return DatedSortable{}
	}
	clientObjects := make(DatedSortable, len(list))
	for i, readObject := range list {
		clientObject := ClientTagEvent{fmt.Sprintf("tag-%d", readObject.Id), "tag", readObject.Claimed, *readObject}
		clientObjects[i] = clientObject
	}

	return clientObjects
}

func (c *Events) GetPlayerEvents(ids_string string) DatedSortable {
	c.Auth()
	query := `
		SELECT player.*
		FROM player
		INNER JOIN event_player
			ON event_player.player_id = player.id
		WHERE event_player.id IN (` + ids_string + `)
    `
	var list []*models.Player
	_, err := Dbm.Select(&list, query)
	if err != nil {
		return DatedSortable{}
	}
	clientObjects := make(DatedSortable, len(list))
	for i, readObject := range list {
		playerRead := PlayerRead{*readObject, readObject.Status(), ""}
		clientObject := ClientPlayerEvent{fmt.Sprintf("joined-%d", readObject.Id), "joined", readObject.Created, playerRead.Id}
		clientObjects[i] = clientObject
	}

	return clientObjects
}

/////////////////////////////////////////////////////////////////////

func addWhereClause(old_where string, to_add string) string {
	if old_where == "" {
		return "WHERE " + to_add
	}
	return old_where + "AND " + to_add
}

func (c *Events) ReadEvents(player_ids []int, game_ids []int) DatedSortable {
	inner_join := ""
	where := ""
	if len(player_ids) > 0 {
		inner_join += "INNER JOIN event_to_player ON event.id = event_to_player.event_id "
		where = addWhereClause(where, "event_to_player.player_id = ANY('{"+IntArrayToString(player_ids)+"}') ")
	}
	if len(game_ids) > 0 {
		inner_join += "INNER JOIN event_to_game ON event.id = event_to_game.event_id "
		where = addWhereClause(where, "event_to_game.game_id = ANY('{"+IntArrayToString(game_ids)+"}') ")
	}

	// don't allow unfiltered events access
	if where == "" {
		return DatedSortable{}
	}

	query := `
		SELECT event.id
		FROM event
    ` + inner_join + where

	var list []*IdWrapper
	_, err := Dbm.Select(&list, query)
	if err != nil {
		panic(err)
	}
	event_ids := make([]int, len(list))
	for i, id := range list {
		event_ids[i] = id.Id
	}
	list = nil
	event_ids_str := IntArrayToString(event_ids)

	var events DatedSortable
	event_ids = nil

	events = append(events, c.GetTagEvents(event_ids_str)...)
	events = append(events, c.GetPlayerEvents(event_ids_str)...)

	sort.Sort(events)

	return events
}

func (c *Events) ReadAllEvents(player_ids []int, game_ids []int) revel.Result {
	events := c.ReadEvents(player_ids, game_ids)
	out := make(map[string]interface{})
	out["events"] = events
	return c.RenderJson(out)
}
