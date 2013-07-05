package tests

import (
	u "flesh/testutils"
	"github.com/robfig/revel"
)

type Generator struct {
	revel.TestSuite
}

func (t *Generator) TestGenerateData() {
	skip := revel.Config.BoolDefault("test.skip_generator", false)
	if skip {
		revel.INFO.Print("Generator disabled, skipping")
		return
	}
	u.GenerateTestData()
}
