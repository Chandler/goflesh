package tests

import (
	"github.com/robfig/revel"
	"strings"
)

type OrganizationTest struct {
	revel.TestSuite
}

// generate some number of organization objects in JSON
func generateOrganizationJson() string {
	jsn := GenerateJson(
		"organizations",
		map[string]func() interface{}{
			"name":             GenerateWord,
			"slug":             GenerateSlug,
			"default_timezone": func() interface{} { return "US/Pacific" },
		},
		-1,
	)

	return jsn
}

func (t OrganizationTest) Before() {
	TestInit()
}

func (t OrganizationTest) TestCreateWorks() {
	orgs := generateOrganizationJson()
	revel.WARN.Print(orgs)
	t.Post("/organizations/", JSON_CONTENT, strings.NewReader(orgs))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	revel.WARN.Print(body)
	t.Assert(strings.Index(body, "default_timezone") != -1)
}

func (t OrganizationTest) TestListWorks() {
	t.Get("/organizations/")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t OrganizationTest) After() {
	TestClean()
}
