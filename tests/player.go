package tests

import (
	"flesh/app/routes"
	u "flesh/testutils"
	sjs "github.com/bitly/go-simplejson"
	"strings"
)

type PlayerTest struct {
	FleshTest
}

// generate some number of organization objects in JSON
func generatePlayerJson() string {
	u.InsertTestUser()
	u.InsertTestOrganization()
	u.InsertTestGame()
	jsn := u.GenerateJson(
		"players",
		map[string]func() interface{}{
			"user_id": func() interface{} { return u.SelectTestUser().Id },
			"game_id": func() interface{} { return u.SelectTestGame().Id },
		},
		1,
	)

	return jsn
}

func (t *PlayerTest) TestCreateAndRead() {
	// create
	jsn := generatePlayerJson()
	t.Post(routes.Players.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "game_id") != -1)

	// read
	responseJson, err := sjs.NewJson(t.ResponseBody)
	t.Assert(err == nil)
	id, err := responseJson.GetIndex(0).Get("id").Int()
	t.Assert(err == nil)
	t.Get(routes.Players.Read(id))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body = string(t.ResponseBody)
	t.Assert(strings.Index(body, "game_id") != -1)

}

func (t *PlayerTest) TestList() {
	t.Get(routes.Players.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}
