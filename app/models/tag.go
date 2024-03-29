package models

import (
	"errors"
	"fmt"
	"time"
)

type Tag struct {
	Id        int        `json:"id"`
	Tagger_id int        `json:"tagger_id"`
	Taggee_id int        `json:"taggee_id"`
	Claimed   *time.Time `json:"claimed"`
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

func NewTag(game *Game, tagger *Player, taggee *Player, claimed *time.Time) (*Tag, int, error) {
	/*
		Conditions for tagging:
			1. Tagger is a zombie (OZ or regular)
			2. Taggee is a human
			3. Game is running
			4. Tagger/taggee in same game

	*/
	if tagger.Game_id != taggee.Game_id {
		return nil, 400, errors.New("Tagger and taggee must be in same game")
	}

	if game == nil {
		gameObj, err := Dbm.Get(Game{}, tagger.Game_id)
		if err != nil {
			return nil, 400, errors.New(fmt.Sprintf("Could not get game with id %d ", tagger.Game_id))
		}
		game = gameObj.(*Game)
	}

	if !game.IsRunning() {
		return nil, 400, errors.New("Game must be running")
	}

	if can := tagger.IsZombie(); !can {
		// it should be ok to print tagger.Status() here because we wouldn't be here if we weren't a zombie or OZ
		return nil, 400, errors.New("Tagger is not zombie, tagger is " + tagger.Status())
	}

	if can := taggee.IsHuman(); !can {
		// it should be ok to print taggee.Status() here because we wouldn't be here if we weren't a zombie or OZ
		return nil, 400, errors.New("Taggee is not human, taggee is " + taggee.Status())
	}

	tag := Tag{0, tagger.Id, taggee.Id, claimed, TimeTrackedModel{}}

	if err := Dbm.Insert(&tag); err != nil {
		return nil, 500, err
	}

	tagger.Feed(claimed)
	tagger.Save()

	taggee.Feed(claimed)
	taggee.Save()

	if err := CreateTagEvent(&tag); err != nil {
		return &tag, 500, err
	}

	return &tag, 201, nil
}
