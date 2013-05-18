package tests

import (
	"flesh/app/models"
	"github.com/coopernurse/gorp"
	r "github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
)

var (
	dbm      *gorp.DbMap
	testMode bool
)

func TestInit() {
	if dbm != nil {
		return
	}

	dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}
	dbm.AddTable(models.Organization{}).SetKeys(true, "Id")
	dbm.AddTable(models.Game{}).SetKeys(true, "Id")
	dbm.AddTable(models.User{}).SetKeys(true, "Id")
	dbm.TraceOn("\x1b[36m[gorp]\x1b[0m", r.INFO)

	testMode, _ := r.Config.Bool("mode.test")
	_ = testMode // used by other testing modules
}
