package models

import (
	"errors"
	"fmt"
)

type Player struct {
	Id      int `json:"id"`
	User_id int `json:"user_id"`
	Game_id int `json:"game_id"`
	TimeTrackedModel
}

func PlayerFromId(id int) (*Player, error) {
	player, err := Dbm.Get(Player{}, id)
	return player.(*Player), err
}

func (p *Player) isZombie() bool {
	fmt.Println(Dbm)
	return true
}

func (p *Player) isHuman() bool {
	return !p.isZombie()
}

func (p *Player) CanTag() (bool, error) {
	isZ := p.isZombie()
	if !isZ {
		return false, errors.New("player is not a zombie!")
	}
	return true, nil
}

func (p *Player) CanBeTagged() (bool, error) {
	isH := p.isHuman()
	if !isH {
		return false, errors.New("player is not a human!")
	}
	return true, nil
}
