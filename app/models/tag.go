package models

import (
	"errors"
	"time"
)

type Tag struct {
	Id        int        `json:"id"`
	Tagger_id int        `json:"tagger_id"`
	Taggee_id int        `json:"taggee_id"`
	Claimed   *time.Time `json:"running_end_time"`
	TimeTrackedModel
}

func NewTag(game *Game, tagger *Player, taggee *Player, claimed *time.Time) (*Tag, error) {
	/*
		Conditions for tagging:
			1. Tagger is a zombie (OZ or regular)
			2. Taggee is a human
			3. Game is running
			4. Tagger/tagee in same game

	*/
	if tagger.Game_id != taggee.Game_id {
		return nil, errors.New("Tagger and tagee must be in same game")
	}

	if !game.IsRunning() {
		return nil, errors.New("Game must be running")
	}

	if can, err := tagger.CanTag(); !can {
		return nil, errors.New("Tagger cannot tag: " + err.Error())
	}

	if can, err := taggee.CanBeTagged(); !can {
		return nil, errors.New("Taggee cannot be tagged: " + err.Error())
	}

	tag := Tag{0, tagger.Id, taggee.Id, claimed, TimeTrackedModel{}}

	return &tag, nil
}
