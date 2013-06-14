package tests

import (
	"flesh/app/controllers"
	"flesh/app/models"
	"flesh/app/routes"
	"github.com/robfig/revel"
	"strings"
	"time"
)

type GameTest struct {
	revel.TestSuite
}

func getOrganizationId() interface{} {
	query := `
    SELECT *
    FROM "organization"
    LIMIT 1
    `
	organizations, _ := controllers.Dbm.Select(models.Organization{}, query)
	organization := organizations[0].(*models.Organization)
	return organization.Id
}

// generate some number of user objects in JSON
func generateGameJson() string {
	testOrg := models.Organization{0, "test org", "test_org", "US/Pacific", nil, nil}
	err := controllers.Dbm.Insert(&testOrg)
	if err != nil {
		panic(err)
	}
	now := time.Now().UTC()
	later := now.Add(12 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)
	tomorrowLater := later.Add(24 * time.Hour)
	jsn := GenerateJson(
		"games",
		map[string]func() interface{}{
			"name":                    GenerateWord,
			"slug":                    GenerateSlug,
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

func (t GameTest) Before() {
	TestInit()
}

func (t GameTest) TestCreateWorks() {
	jsn := generateGameJson()
	t.Post(routes.Games.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "registration_start_time") != -1)
}

func (t GameTest) TestListWorks() {
	t.Get(routes.Games.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}

func (t GameTest) After() {
	TestClean()
}
