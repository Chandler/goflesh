package tests

import (
	"flesh/app/models"
	"github.com/robfig/revel"
	"net/url"
	"strings"
)

type PlayerTest struct {
	revel.TestSuite
}

func getUserId() interface{} {
	query := `
    SELECT *
    FROM "user"
    LIMIT 1
    `
	users, _ := dbm.Select(models.User{}, query)
	user := users[0].(*models.User)
	return user.Id
}

func getGameId() interface{} {
	query := `
    SELECT *
    FROM "game"
    LIMIT 1
    `
	games, _ := dbm.Select(models.Game{}, query)
	game := games[0].(*models.Game)
	return game.Id
}

// generate some number of organization objects in JSON
func generatePlayerJson() string {
	jsn := GenerateJson(
		map[string]func() interface{}{
			"user_id": getUserId,
			"game_id": getGameId,
		},
		-1,
	)

	return jsn
}

func (t PlayerTest) Before() {
	TestInit()
}

func (t PlayerTest) TestCreateWorks() {
	orgs := url.Values{}
	orgs.Add("data", generatePlayerJson())
	t.PostForm("/players/", orgs)
	t.AssertOk()
	t.AssertContentType("application/json")
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "game_id") != -1)
}

func (t PlayerTest) TestListWorks() {
	t.Get("/players/")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t PlayerTest) After() {
	if testMode {
		dbm.Exec("TRUNCATE player CASCADE")
	}
}
