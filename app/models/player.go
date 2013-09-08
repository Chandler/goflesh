package models

import (
	"errors"
	"github.com/robfig/revel"
	"time"
)

type Player struct {
	Id       int        `json:"id"`
	User_id  int        `json:"user_id"`
	Game_id  int        `json:"game_id"`
	Last_fed *time.Time `json:"last_fed"`
	TimeTrackedModel
}

func (p *Player) HumanCode() *HumanCode {
	human, err := Dbm.Get(HumanCode{}, p.Id)
	if err != nil {
		revel.ERROR.Print("Could not get human code", err)
	}
	return human.(*HumanCode)
}

func (p *Player) User() *User {
	user, err := UserFromId(p.User_id)
	if err != nil {
		panic(err)
	}
	return user
}

func PlayerFromId(id int) (*Player, error) {
	player, err := Dbm.Get(Player{}, id)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, errors.New("Player could not be found")
	}
	return player.(*Player), err
}

func PlayerFromHumanCode(code string) (*Player, int, error) {
	query := `
        SELECT *
        FROM human_code
        WHERE code = $1
    `
	var list []*HumanCode
	_, err := Dbm.Select(&list, query, code)
	if err != nil {
		return nil, 500, err
	}
	if len(list) != 1 {
		return nil, 422, errors.New("Invalid human code")
	}
	human_code := list[0]
	player, err := Dbm.Get(Player{}, human_code.Id)
	if err != nil {
		return nil, 500, err
	}

	human := player.(*Player)
	return human, 200, nil
}

func (m *Player) Save() error {
	_, err := Dbm.Update(m)
	return err
}

func (p *Player) Game() *Game {
	game, err := GameFromId(p.Game_id)
	if err != nil {
		panic(err)
	}
	return game
}

func PlayerFromUserIdGameId(user_id int, game_id int) (*Player, error) {
	query := `
		SELECT *
		FROM player
		WHERE user_id = $1
		AND game_id = $2
	`

	var list []*Player
	_, err := Dbm.Select(&list, query)
	if err != nil {
		return nil, err
	}

	if len(list) != 1 {
		return nil, errors.New("Could not get player object")
	}

	return list[0], err
}

func (p *Player) Feed(fedAt *time.Time) {
	if fedAt == nil {
		now := time.Now()
		fedAt = &now
	}
	p.Last_fed = fedAt
}

func (p *Player) Status() string {
	if p.Last_fed == nil {
		return "human"
	}

	// TODO: support custom TimeToStarve without looking up game objects in DB
	// mustHaveFedBy := time.Now().Add(-p.Game().TimeToStarve())
	mustHaveFedBy := time.Now().Add(-new(Game).TimeToStarve())

	if p.Last_fed.Before(mustHaveFedBy) {
		return "starved"
	}

	return "zombie"
}

func (p *Player) CalculateStatus() string {
	query := `
		WITH zombies AS (
			SELECT p.id
			FROM player p
			LEFT OUTER JOIN "oz"
			        ON p.id = oz.id
			LEFT OUTER JOIN "tag"
			        ON p.id = taggee_id
			WHERE p.id = $1
			        AND (oz.id IS NULL OR oz.confirmed = FALSE)
			        AND taggee_id IS NULL
		)
		SELECT CASE
			WHEN count(zombies.id) = 0 THEN 'human'
			WHEN EXISTS (
				SELECT tag.id
				FROM tag
				INNER JOIN zombies z
				ON tag.tagger_id = z.id
				AND claimed > $2
				) THEN 'zombie'
			ELSE 'starved'
		END
		FROM zombies
	`
	game, err := GameFromId(p.Game_id)
	if err != nil {
		panic(err)
	}

	mustTagAfter := time.Now().Add(-game.TimeToStarve())

	status, err := Dbm.SelectStr(query, p.Id, mustTagAfter)
	if err != nil {
		panic(err)
	}

	return status
}

func (p *Player) TaggedTag() *Tag {
	query := `
		SELECT *
		FROM "tag"
		WHERE taggee_id = $1
	`
	var list []*Tag
	_, err := Dbm.Select(&list, query)
	if err != nil {
		panic(err)
	}
	return list[0]
}

func (p *Player) GetLastTag() *Tag {
	query := `
		SELECT *
		FROM "tag"
		WHERE tagger_id = $1
		ORDER BY claimed DESC
		LIMIT 1
	`
	var list []*Tag
	_, err := Dbm.Select(&list, query, p.Id)
	if err != nil {
		panic(err)
	}

	if len(list) == 0 {
		// this player has never tagged anybody
		return nil
	}
	return list[0]
}

func (p *Player) IsStarved() bool {
	return p.Status() == "starved"
}

func (p *Player) IsZombie() bool {
	return p.Status() == "zombie"
}

func (p *Player) IsHuman() bool {
	return p.Status() == "human"
}
