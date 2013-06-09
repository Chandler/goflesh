package tests

import (
	"github.com/robfig/revel"
	"net/url"
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
	orgs := url.Values{}
	orgs.Add("organizations", generateOrganizationJson())
	revel.INFO.Print(orgs)
	t.PostForm("/organizations/", orgs)
	t.AssertOk()
	t.AssertContentType("application/json")
	body := string(t.ResponseBody)
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
