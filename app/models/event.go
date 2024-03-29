package models

import "fmt"

var (
	// TODO: fetch these values from the database at startup instead of hard-coding them
	// EventRole ID values
	EVENT_ROLE_TAGGER_VALUE = 1
	EVENT_ROLE_TAGGEE_VALUE = 2
	EVENT_ROLE_JOINER_VALUE = 3
	// EventRole IDs
	EVENT_ROLE_TAGGER = &EVENT_ROLE_TAGGER_VALUE
	EVENT_ROLE_TAGGEE = &EVENT_ROLE_TAGGEE_VALUE
	EVENT_ROLE_JOINER = &EVENT_ROLE_JOINER_VALUE
)

// An Event is an occurrence that can be displayed in an event feed
type Event struct {
	Id int `json:"id"`
	TimeTrackedModel
}

// They type of an event
type EventType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Table_name  string `json:"-,omitempty"`
}

// The Role an object plays in an Event -- perhaps a "tagger" player in a TagEvent
type EventRole struct {
	Id            int    `json:"id"`
	Event_type_id int    `json:"event_type"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

// M2M for a Player participating in an Event
type EventToPlayer struct {
	Id            int  `json:"id"`
	Event_id      int  `json:"event"`
	Player_id     int  `json:"player"`
	Event_role_id *int `json:"event_role"`
}

// M2M for a Game participating in an Event
type EventToGame struct {
	Id       int `json:"id"`
	Event_id int `json:"event"`
	Game_id  int `json:"game"`
}

// An Event representing a Tag
type EventTag struct {
	Id     int `json:"id"` // FK to Event
	Tag_id int `json:"tag"`
}

// An Event representing a Player creation (joining a game)
type EventPlayer struct {
	Id        int `json:"id"` // FK to Event
	Player_id int `json:"player"`
}

func CreateTagEvent(tag *Tag) error {
	// create the base event
	event := &Event{0, TimeTrackedModel{}}
	err := Dbm.Insert(event)
	if err != nil {
		return err
	}

	fmt.Println("Creating tag event. Event id: %v", event.Id)

	// create the Tag event
	tagEvent := &EventTag{event.Id, tag.Id}
	err = Dbm.Insert(tagEvent)
	if err != nil {
		return err
	}

	// record players involved in this event
	tagger_m2m := &EventToPlayer{0, event.Id, tag.Tagger_id, EVENT_ROLE_TAGGER}
	err = Dbm.Insert(tagger_m2m)
	if err != nil {
		return err
	}
	taggee_m2m := &EventToPlayer{0, event.Id, tag.Taggee_id, EVENT_ROLE_TAGGEE}
	err = Dbm.Insert(taggee_m2m)
	if err != nil {
		return err
	}

	// record game associated with this event
	game_m2m := &EventToGame{0, event.Id, tag.Tagger().Game_id}
	err = Dbm.Insert(game_m2m)
	if err != nil {
		return err
	}

	return nil
}

func CreateJoinedGameEvent(player *Player) error {
	// create the base event
	event := Event{0, TimeTrackedModel{}}
	err := Dbm.Insert(&event)
	if err != nil {
		return err
	}

	// create the Tag event
	playerEvent := EventPlayer{event.Id, player.Id}
	err = Dbm.Insert(&playerEvent)
	if err != nil {
		return err
	}

	// record players involved in this event
	player_m2m := EventToPlayer{0, event.Id, player.Id, EVENT_ROLE_JOINER}
	err = Dbm.Insert(&player_m2m)
	if err != nil {
		return err
	}

	// record game associated with this event
	game_m2m := EventToGame{0, event.Id, player.Game_id}
	err = Dbm.Insert(&game_m2m)
	if err != nil {
		return err
	}

	return nil
}
