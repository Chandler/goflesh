package tests

import (
	"flesh/app/models"
	"github.com/robfig/revel"
	"net/url"
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
	organizations, _ := dbm.Select(models.Organization{}, query)
	organization := organizations[0].(*models.Organization)
	return organization.Id
}

// generate some number of user objects in JSON
func generateGameJson() string {
	now := time.Now().UTC()
	later := now.Add(12 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)
	tomorrowLater := later.Add(24 * time.Hour)
	jsn := GenerateJson(
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
	orgs := url.Values{}
	orgs.Add("data", generateGameJson())
	t.PostForm("/games/", orgs)
	t.AssertOk()
	t.AssertContentType("application/json")
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "registration_start_time") != -1)
}

func (t GameTest) TestListWorks() {
	t.Get("/games/")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t GameTest) After() {
	TestClean()
}
