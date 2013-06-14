package tests

import (
	"flesh/app/routes"
	"github.com/robfig/revel"
	"strings"
)

type UserTest struct {
	revel.TestSuite
}

// generate some number of user objects in JSON
func generateUserJson() string {
	jsn := GenerateJson(
		"users",
		map[string]func() interface{}{
			"email":       GenerateEmail,
			"first_name":  GenerateWord,
			"last_name":   GenerateWord,
			"screen_name": GenerateSlug,
			"password":    func() interface{} { return GenerateString(4, "-X-") },
		},
		-1,
	)

	return jsn
}

func (t UserTest) Before() {
	TestInit()
}

func (t UserTest) TestCreateWorks() {
	jsn := generateUserJson()
	t.Post(routes.Users.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "first_name") != -1)
}

func (t UserTest) TestListWorks() {
	t.Get(routes.Users.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}

func (t UserTest) After() {
	TestClean()
}
