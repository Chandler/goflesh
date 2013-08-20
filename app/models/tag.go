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
		return nil, errors.New("Tagger and taggee must be in same game")
	}

	if game == nil {
		gameObj, err := Dbm.Get(Game{}, tagger.Game_id)
		if err != nil {
			return nil, errors.New("Could not get game")
		}
		game = gameObj.(*Game)
	}

	if !game.IsRunning() {
		return nil, errors.New("Game must be running")
	}

	if can := tagger.IsZombie(); !can {
		return nil, errors.New("Tagger is not zombie")
	}

	if can := taggee.IsHuman(); !can {
		return nil, errors.New("Taggee is not human")
	}

	tag := Tag{0, tagger.Id, taggee.Id, claimed, TimeTrackedModel{}}

	if err := Dbm.Insert(&tag); err != nil {
		return nil, err
	}

	CreateTagEvent(&tag)

	return &tag, nil
}
