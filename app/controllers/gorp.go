package controllers

import (
	"database/sql"
	// "github.com/bmizerany/pq"
	// "code.google.com/p/go.crypto/bcrypt"
	"flesh/app/models"
	"github.com/coopernurse/gorp"
	r "github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
)

var (
	dbm *gorp.DbMap
)

type GorpPlugin struct {
	r.EmptyPlugin
}

func (p GorpPlugin) OnAppStart() {
	db.DbPlugin{}.OnAppStart()
	dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	dbm.AddTable(models.Organization{}).SetKeys(true, "Id")
	dbm.TraceOn("\x1b[36m[gorp]\x1b[0m", r.INFO)

	// Create tables (ok if they exist, move on)
	// TODO: replace this with official schema
	dbm.CreateTables()

	organizations := []*models.Organization{
		&models.Organization{0, "UIdaho", "vandals"},
		&models.Organization{0, "Boise State", "broncs"},
		&models.Organization{0, "Berkeley", "cal"},
	}

	for _, organization := range organizations {
		if err := dbm.Insert(organization); err != nil {
			// panic(err)
		}
	}
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
