package tests

import (
	"flesh/app/routes"
	"github.com/robfig/revel"
	"strings"
)

type PlayerTest struct {
	revel.TestSuite
}

// generate some number of organization objects in JSON
func generatePlayerJson() string {
	InsertTestUser()
	InsertTestOrganization()
	InsertTestGame()
	jsn := GenerateJson(
		"players",
		map[string]func() interface{}{
			"user_id": func() interface{} { return SelectTestUser().Id },
			"game_id": func() interface{} { return SelectTestGame().Id },
		},
		1,
	)

	return jsn
}

func (t PlayerTest) Before() {
	TestInit()
}

func (t PlayerTest) TestCreateWorks() {
	jsn := generatePlayerJson()
	t.Post(routes.Players.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "game_id") != -1)
}

func (t PlayerTest) TestListWorks() {
	t.Get(routes.Games.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}

func (t PlayerTest) After() {
	TestClean()
}
