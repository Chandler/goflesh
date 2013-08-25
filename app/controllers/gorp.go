package controllers

import (
	"database/sql"
	"flesh/app/models"
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	"math/rand"
	"time"
)

var (
	Dbm *gorp.DbMap
)

func GorpInit() {
	rand.Seed(time.Now().UTC().UnixNano())
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}
	Dbm.TraceOn("\x1b[36m[C.gorp]\x1b[0m", revel.INFO)
	models.AddTables(Dbm)
}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) isDevMode() bool {
	return revel.Config.BoolDefault("mode.dev", false)
}
