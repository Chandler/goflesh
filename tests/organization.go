package tests

import (
	"flesh/app/routes"
	u "flesh/testutils"
	sjs "github.com/bitly/go-simplejson"
	"strings"
)

type OrganizationTest struct {
	FleshTest
}

// generate some number of organization objects in JSON
func generateOrganizationJson() string {
	jsn := u.GenerateJson(
		"organizations",
		map[string]func() interface{}{
			"name":             u.GenerateWord,
			"slug":             u.GenerateSlug,
			"location":         u.GenerateName,
			"default_timezone": func() interface{} { return "US/Pacific" },
		},
		-1,
	)

	return jsn
}

func (t *OrganizationTest) TestCreateAndRead() {
	// create
	jsn := generateOrganizationJson()
	t.Post(routes.Organizations.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "default_timezone") != -1)

	// read
	responseJson, err := sjs.NewJson(t.ResponseBody)
	t.Assert(err == nil)
	id, err := responseJson.GetIndex(0).Get("id").Int()
	t.Assert(err == nil)
	t.Get(routes.Organizations.Read(id))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body = string(t.ResponseBody)
	t.Assert(strings.Index(body, "default_timezone") != -1)
}

func (t *OrganizationTest) TestList() {
	t.Get(routes.Organizations.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}
