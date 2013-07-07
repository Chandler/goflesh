package tests

import (
	"flesh/app/controllers"
	"flesh/app/models"
	"flesh/app/routes"
	u "flesh/testutils"
	sjs "github.com/bitly/go-simplejson"
	"strings"
	"time"
)

type GameTest struct {
	FleshTest
}

func getOrganizationId() interface{} {
	organization := u.SelectTestOrganization()
	return organization.Id
}

// generate some number of user objects in JSON
func generateGameJson() string {
	u.InsertTestOrganization()
	now := time.Now().UTC()
	later := now.Add(12 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)
	tomorrowLater := later.Add(24 * time.Hour)
	jsn := u.GenerateJson(
		"games",
		map[string]func() interface{}{
			"name":                    u.GenerateWord,
			"slug":                    u.GenerateSlug,
			"organization_id":         getOrganizationId,
			"timezone":                func() interface{} { return "US/Pacific" },
			"registration_start_time": func() interface{} { return now.Format(time.RFC3339) },
			"registration_end_time":   func() interface{} { return tomorrow.Format(time.RFC3339) },
			"running_start_time":      func() interface{} { return later.Format(time.RFC3339) },
			"running_end_time":        func() interface{} { return tomorrowLater.Format(time.RFC3339) },
		},
		1,
	)

	return jsn
}

func (t *GameTest) TestCreateAndRead() {
	// create
	jsn := generateGameJson()
	t.Post(routes.Games.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "registration_start_time") != -1)

	// read
	responseJson, err := sjs.NewJson(t.ResponseBody)
	t.Assert(err == nil)
	id, err := responseJson.GetIndex(0).Get("id").Int()
	t.Assert(err == nil)
	t.Get(routes.Games.Read(id))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body = string(t.ResponseBody)

	t.Assert(strings.Index(body, "registration_start_time") != -1)
}

func (t *GameTest) TestList() {
	t.Get(routes.Games.ReadList([]int{}))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}

func (t *GameTest) TestIsRunning() {
	u.InsertTestOrganization()
	org := u.SelectTestOrganization()
	now := time.Now()
	twoDaysAgo := now.Add(u.TwoDaysBack)
	twoDaysHence := now.Add(u.TwoDaysForward)
	oneDayAgo := now.Add(u.OneDayBack)
	oneDayHence := now.Add(u.OneDayForward)
	current := &models.Game{0,
		"name",
		"slug",
		org.Id,
		"US/Pacific",
		&twoDaysAgo,
		&oneDayHence,
		&oneDayAgo,
		&twoDaysHence,
		models.TimeTrackedModel{},
	}
	err := controllers.Dbm.Insert(current)
	t.Assert(err == nil)
	t.Assert(current.IsRunning())

	past := &models.Game{0,
		"name",
		"slug",
		org.Id,
		"US/Pacific",
		&twoDaysAgo,
		&oneDayHence,
		&twoDaysAgo,
		&oneDayAgo,
		models.TimeTrackedModel{},
	}
	err = controllers.Dbm.Insert(past)
	t.Assert(err == nil)
	t.Assert(!past.IsRunning())

	future := &models.Game{0,
		"name",
		"slug",
		org.Id,
		"US/Pacific",
		&twoDaysAgo,
		&oneDayHence,
		&oneDayHence,
		&twoDaysHence,
		models.TimeTrackedModel{},
	}
	err = controllers.Dbm.Insert(future)
	t.Assert(err == nil)
	t.Assert(!future.IsRunning())
}
