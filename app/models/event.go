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

type EventTag struct {
	Id          int `json:"id"`
	EventId     int `json:"event"`
	TagId       int `json:"tag"`
	EventRoleId int `json:"event_role"`
}

func CreateTagEvent(tag *Tag) error {
	// TODO: for heaven's sake, don't hardcode these ids!
	event := Event{0, 1, &TimeTrackedModel{}}
	err := Dbm.Insert(event)
	if err != nil {
		return err
	}
	tagEvent := EventTag{0, event.Id, tag.Id, 1}
	err = Dbm.Insert(tagEvent)
	if err != nil {
		return err
	}
	return nil
}
