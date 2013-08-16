package models

import ()

type Event struct {
	Id          int `json:"id"`
	EventTypeId int `json:"event_type"`
	*TimeTrackedModel
}

type EventType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EventRole struct {
	Id          int    `json:"id"`
	EventTypeId int    `json:"event_type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EventPlayer struct {
	Id          int `json:"id"`
	EventId     int `json:"event"`
	PlayerId    int `json:"player"`
	EventRoleId int `json:"event_role"`
}

func CreateTagEvent(tag *Tag) error {
	// TODO: for heaven's sake, don't hardcode these ids!
	event := Event{0, 1, &TimeTrackedModel{}}
	err := Dbm.Insert(event)
	if err != nil {
		return err
	}
	tagger := EventPlayer{0, event.Id, tag.Tagger_id, 1}
	err = Dbm.Insert(tagger)
	if err != nil {
		return err
	}
	taggee := EventPlayer{0, event.Id, tag.Taggee_id, 2}
	err = Dbm.Insert(taggee)
	if err != nil {
		return err
	}
	return nil
}
