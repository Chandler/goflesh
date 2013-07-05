package tests

import (
	"github.com/robfig/revel"
)

const (
	JSON_CONTENT string = "application/json"
)

type ApplicationTest struct {
	revel.TestSuite
}

type FleshTest struct {
	revel.TestSuite
}

func (t *FleshTest) Before() {
	TestInit()
}

func (t *FleshTest) After() {
	TestClean()
}

func (t *ApplicationTest) Before() {
	TestInit()
}

func (t *ApplicationTest) TestIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t *ApplicationTest) After() {
	TestClean()
}
