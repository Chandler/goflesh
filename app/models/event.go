package models

const (
	// TODO: fetch these values from the database at startup instead of hard-coding them
	// EventType IDs
	EVENT_TYPE_TAG = 1
)

var (
	// TODO: fetch these values from the database at startup instead of hard-coding them
	// EventRole ID values
	EVENT_ROLE_TAGGER_VALUE = 1
	EVENT_ROLE_TAGGEE_VALUE = 2
	// EventRole IDs
	EVENT_ROLE_TAGGER = &EVENT_ROLE_TAGGER_VALUE
	EVENT_ROLE_TAGGEE = &EVENT_ROLE_TAGGEE_VALUE
)

// An Event is an occurrence that can be displayed in an event feed
type Event struct {
	Id            int `json:"id"`
	Event_type_id int `json:"event_type"`
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

func CreateTagEvent(tag *Tag) error {
	// create the base event
	event := Event{0, 1, TimeTrackedModel{}} // TODO: remove hardcoded id
	err := Dbm.Insert(&event)
	if err != nil {
		return err
	}

	// create the Tag event
	tagEvent := EventTag{event.Id, tag.Id}
	err = Dbm.Insert(&tagEvent)
	if err != nil {
		return err
	}

	// record players involved in this event
	tagger_m2m := EventToPlayer{0, event.Id, tag.Tagger_id, EVENT_ROLE_TAGGER}
	err = Dbm.Insert(&tagger_m2m)
	if err != nil {
		return err
	}
	taggee_m2m := EventToPlayer{0, event.Id, tag.Taggee_id, EVENT_ROLE_TAGGEE}
	err = Dbm.Insert(&taggee_m2m)
	if err != nil {
		return err
	}

	// record game associated with this event
	game_m2m := EventToGame{0, event.Id, tag.Tagger().Game_id}
	err = Dbm.Insert(&game_m2m)
	if err != nil {
		return err
	}

	return nil
}
