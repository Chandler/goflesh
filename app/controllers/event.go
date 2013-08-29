package controllers

import (
	"flesh/app/models"
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
	return s[i].Date().Before(*s[j].Date())
}
func (s DatedSortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ClientTagEvent struct {
	Type     string     `json:"type"`
	SortDate *time.Time `json:"-"`
	Tag      models.Tag `json:"tag"`
}

func (c ClientTagEvent) Date() *time.Time {
	return c.SortDate
}

func (c *Events) GetTagEvents(inner_join string, where string, args ...interface{}) DatedSortable {
	query := `
		SELECT tag.*
		FROM event
		` + inner_join + `
		INNER JOIN event_tag
			ON event.id = event_tag.id
		INNER JOIN tag
			ON event_tag.tag_id = tag.id
    ` + where + `
    	ORDER BY tag.claimed DESC
    `
	var list []*models.Tag
	_, err := Dbm.Select(&list, query, args...)
	if err != nil {
		return DatedSortable{}
	}
	clientObjects := make(DatedSortable, len(list))
	for i, readObject := range list {
		clientObject := ClientTagEvent{"tag", readObject.Claimed, *readObject}
		clientObjects[i] = clientObject
	}

	sort.Sort(clientObjects)

	return clientObjects
}

/////////////////////////////////////////////////////////////////////

func (c *Events) ReadPlayers(ids []int) revel.Result {
	if len(ids) == 0 {
		return c.RenderText("")
	}
	templateStr := IntArrayToString(ids)
	events := c.GetTagEvents(
		"INNER JOIN event_to_player ON event.id = event_to_player.event_id",
		"WHERE event_to_player.id = ANY('{"+templateStr+"}')",
	)
	return c.RenderJson(events)
}

// func (c *Games) ReadList(ids []int) revel.Result {
// 	if len(ids) == 0 {
// 		return c.ReadGame("")
// 	}
// 	templateStr := IntArrayToString(ids)
// 	return c.ReadGame("WHERE g.id = ANY('{" + templateStr + "}')")
// }

// /////////////////////////////////////////////////////////////////////

// func (c *Games) Read(id int) revel.Result {
// 	return c.ReadGame("WHERE g.id = $1", id)
// }
