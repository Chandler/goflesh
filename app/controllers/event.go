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
	Id        string     `json:"id"`
	Type      string     `json:"type"`
	SortDate  *time.Time `json:"-"`
	Player_id int        `json:"player_id"`
}

func (c ClientPlayerEvent) Date() *time.Time {
	return c.SortDate
}

type IdWrapper struct {
	Id int
}

type TagEventRead struct {
	models.Tag
	Is_oz *bool
}

func (c *Events) GetTagEvents(ids_string string, filtered_on_player bool, player_ids_to_include_even_if_oz_tagged map[int]bool) DatedSortable {
	query := `
		SELECT tag.*, oz.confirmed is_oz
		FROM tag
		INNER JOIN event_tag
			ON event_tag.tag_id = tag.id
		LEFT JOIN oz
			ON oz.id = tag.tagger_id
		WHERE event_tag.id IN (` + ids_string + `)
    `
	var list []*TagEventRead
	_, err := Dbm.Select(&list, query)
	if err != nil {
		revel.ERROR.Print(err)
		return DatedSortable{}
	}
	clientObjects := make(DatedSortable, 0, len(list))
	for i, readObject := range list {
		// if this player is an OZ, blank out if OZ status should not yet be revealed
		if readObject.Is_oz != nil && *readObject.Is_oz {
			tagger := readObject.Tag.Tagger()
			if tagger.IsZombie() && !tagger.Game().OzsAreRevealed() {
				readObject.Tagger_id = models.OZ_PLAYER_ID
				// if we filtered by player_id and this player wasn't on the receiving end of that tag, don't include it at all
				if filtered_on_player {
					should_include, in_set := player_ids_to_include_even_if_oz_tagged[readObject.Tag.Taggee_id]
					if !in_set || !should_include {
						continue
					}
				}
			}
		}
		clientObject := ClientTagEvent{fmt.Sprintf("tag-%d", readObject.Id), "tag", readObject.Claimed, readObject.Tag}
		clientObjects = clientObjects[:len(clientObjects)+1]
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
	for i, player := range list {
		clientObject := ClientPlayerEvent{fmt.Sprintf("joined-%d", player.Id), "joined", player.Created, player.Id}
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
		// inner_join += "INNER JOIN event_to_game ON event.id = event_to_game.event_id "
		// where = addWhereClause(where, "event_to_game.game_id = ANY('{"+IntArrayToString(game_ids)+"}') ")

		// TODO: fix event_to_game inserts so this isn't necessary
		inner_join += "INNER JOIN event_to_player ON event.id = event_to_player.event_id INNER JOIN player ON event_to_player.player_id = player.id "
		where = addWhereClause(where, "player.game_id = ANY('{"+IntArrayToString(game_ids)+"}') ")
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

	filteredOnPlayer := len(player_ids) > 0
	var include_events_for_player_ids map[int]bool
	if filteredOnPlayer {
		include_events_for_player_ids = make(map[int]bool)
		for _, id := range player_ids {
			include_events_for_player_ids[id] = true
		}
	}
	events = append(events, c.GetTagEvents(event_ids_str, filteredOnPlayer, include_events_for_player_ids)...)
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
