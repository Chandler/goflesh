package tests

import (
	"encoding/json"
	"flesh/app/routes"
	"github.com/robfig/revel"
	"strings"
)

type UserTest struct {
	FleshTest
}

// generate some number of user objects in JSON
func generateUserStructArray() []map[string]interface{} {
	structArray := GenerateStructArray(
		map[string]func() interface{}{
			"email":       GenerateEmail,
			"first_name":  GenerateWord,
			"last_name":   GenerateWord,
			"screen_name": GenerateSlug,
			"password":    func() interface{} { return GenerateString(4, "-X-") },
		},
		-1,
	)

	return structArray
}

func generateUserJson() string {
	embedded := EmbedMapUnderKey("users", generateUserStructArray())
	return ConvertMappedStructArrayToString(embedded)
}

func (t *UserTest) TestCreate() {
	jsn := generateUserJson()
	t.Post(routes.Users.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "first_name") != -1)
}

func (t *UserTest) TestList() {
	t.Get(routes.Users.ReadList())
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)
}

func (t *UserTest) TestAuthenticate() {
	structArray := generateUserStructArray()
	jsn := ConvertMappedStructArrayToString(EmbedMapUnderKey("users", structArray))

	// Create the users
	t.Post(routes.Users.Create(), JSON_CONTENT, strings.NewReader(jsn))
	t.AssertOk()
	t.AssertContentType(JSON_CONTENT)

	// Loop over the users, test with valid and invalid password
	for _, user := range structArray {
		userAuth := map[string]string{
			"email":       user["email"].(string),
			"screen_name": user["screen_name"].(string),
			"password":    user["password"].(string),
		}

		// Test with valid password
		jsn, err := json.Marshal(userAuth)
		if err != nil {
			revel.ERROR.Fatal(err)
		}
		jsonReader := strings.NewReader(string(jsn))
		t.Post(routes.Users.Authenticate(), JSON_CONTENT, jsonReader)
		t.AssertOk()
		t.AssertContentType(JSON_CONTENT)

		// Test with invalid password
		userAuth["password"] = userAuth["password"] + "1"
		jsn, err = json.Marshal(userAuth)
		if err != nil {
			revel.ERROR.Fatal(err)
		}
		jsonReader = strings.NewReader(string(jsn))
		t.Post(routes.Users.Authenticate(), JSON_CONTENT, jsonReader)
		t.AssertStatus(401)
	}
}
