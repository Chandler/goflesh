package tests

import (
	"flesh/app/routes"
	u "flesh/testutils"
	sjs "github.com/bitly/go-simplejson"
	"strings"
)

type MemberTest struct {
	FleshTest
}

// generate some number of organization objects in JSON
func generateMemberJson() string {
	u.InsertTestUser()
	u.InsertTestOrganization()
	u.InsertTestGame()
	jsn := u.GenerateJson(
		"members",
		map[string]func() interface{}{
			"user_id": func() interface{} { return u.SelectTestUser().Id },
			"organization_id": func() interface{} { return u.SelectTestOrganization().Id },
		},
		1,
	)

	return jsn
}

func (t *MemberTest) TestCreate() {
	// create
	jsn := generateMemberJson()
	t.Post(routes.Members.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "organization_id") != -1)
}

func (t *MemberTest) TestRead() {
	t.TestCreate()
	body := string(t.ResponseBody)

	// read
	responseJson, err := sjs.NewJson(t.ResponseBody)
	t.Assert(err == nil)
	id, err := responseJson.GetIndex(0).Get("id").Int()
	t.Assert(err == nil)
	t.Get(routes.Members.Read(id))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body = string(t.ResponseBody)
	t.Assert(strings.Index(body, "organization_id") != -1)
}

func (t *MemberTest) TestList() {
	t.Get(routes.Members.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}
