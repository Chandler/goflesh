package models

type Oz struct {
	Id        int  `json:"id"`
	Confirmed bool `json:"confirmed"`
	TimeTrackedModel
}

func (m *Oz) Confirm() {
	m.Confirmed = true
	err := m.Save()
	if err != nil {
		panic(err)
	}
	player := m.Player()
	game := player.Game()
	player.Feed(game.TimeToReveal())
}

func (m *Oz) Save() error {
	_, err := Dbm.Update(m)
	return err
}

func (m *Oz) Player() *Player {
	player, err := PlayerFromId(m.Id)
	if err != nil {
		panic(err)
	}
	return player
}
