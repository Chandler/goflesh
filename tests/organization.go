package tests

import (
	"fmt"
	"github.com/robfig/revel"
	"net/url"
	"strings"
)

type OrganizationTest struct {
	revel.TestSuite
}

var (
	_                = fmt.Println // for test development
	organizationJson = `[
    {
      "name": "University of Idaho test",
      "slug": "vandals_test",
      "default_timezone": "US/Pacific"
    },
    {
      "name": "Boise State test",
      "slug": "broncs_test",
      "default_timezone": "US/Mountain"
    }
    ]
    `
)

func (t OrganizationTest) Before() {
	TestInit()
}

func (t OrganizationTest) TestCreateWorks() {
	orgs := url.Values{}
	orgs.Add("data", organizationJson)
	fmt.Println(orgs)
	t.PostForm("/organizations/", orgs)
	t.AssertOk()
	t.AssertContentType("application/json")
	body := string(t.ResponseBody)
	t.Assert(strings.Index(body, "University of Idaho") != -1)
	t.Assert(strings.Index(body, "broncs_test") != -1)
}

func (t OrganizationTest) TestListWorks() {
	t.Get("/organizations/")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t OrganizationTest) After() {
	if testMode {
		dbm.Exec("TRUNCATE organization CASCADE")
	}
}
