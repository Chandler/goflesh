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

func NewTag(game *Game, tagger *Player, taggee *Player, claimed *time.Time) error {
	/*
		Conditions for tagging:
			1. Tagger is a zombie (OZ or regular)
			2. Taggee is a human
			3. Game is running
			4. Tagger/tagee in same game

	*/
	if tagger.Game_id != taggee.Game_id {
		return errors.New("Tagger and tagee must be in same game")
	}

	if !game.IsRunning() {
		return errors.New("Game must be running")
	}

	if can, err := tagger.CanTag(); !can {
		return errors.New("Tagger cannot tag: " + err.Error())
	}

	if can, err := tagger.CanBeTagged(); !can {
		return errors.New("Taggee cannot be tagged: " + err.Error())
	}

	tag := &Tag{0, tagger.Id, taggee.Id, claimed, TimeTrackedModel{}}

	err := Dbm.Insert(tag)
	return err
}
