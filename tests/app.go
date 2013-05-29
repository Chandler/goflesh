package tests

import (
	"github.com/robfig/revel"
)

type ApplicationTest struct {
	revel.TestSuite
}

func (t ApplicationTest) Before() {
}

func (t ApplicationTest) TestIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t ApplicationTest) After() {
}
