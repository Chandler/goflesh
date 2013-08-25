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

func (t *Tag) Tagger() *Player {
	player, err := PlayerFromId(t.Tagger_id)
	if err != nil {
		panic(err)
	}
	return player
}

func (t *Tag) Taggee() *Player {
	player, err := PlayerFromId(t.Taggee_id)
	if err != nil {
		panic(err)
	}
	return player
}

func NewTag(game *Game, tagger *Player, taggee *Player, claimed *time.Time) (*Tag, error) {
	/*
		Conditions for tagging:
			1. Tagger is a zombie (OZ or regular)
			2. Taggee is a human
			3. Game is running
			4. Tagger/taggee in same game

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
		// it should be ok to print tagger.Status() here because we wouldn't be here if we weren't a zombie or OZ
		return nil, errors.New("Tagger is not zombie: " + tagger.Status())
	}

	if can := taggee.IsHuman(); !can {
		// it should be ok to print taggee.Status() here because we wouldn't be here if we weren't a zombie or OZ
		return nil, errors.New("Taggee is not human" + taggee.Status())
	}

	tag := Tag{0, tagger.Id, taggee.Id, claimed, TimeTrackedModel{}}

	if err := Dbm.Insert(&tag); err != nil {
		return nil, err
	}

	tagger.Feed(claimed)
	tagger.Save()

	if err := CreateTagEvent(&tag); err != nil {
		return &tag, err
	}

	return &tag, nil
}
