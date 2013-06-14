package tests

import (
	"flesh/app/routes"
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
	jsn := generateOrganizationJson()
	t.Post(routes.Organizations.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "default_timezone") != -1)
}

func (t OrganizationTest) TestListWorks() {
	t.Get(routes.Organizations.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}

func (t OrganizationTest) After() {
	TestClean()
}
