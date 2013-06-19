package controllers

import (
	"database/sql"
	"flesh/app/models"
	"github.com/msolomon/gorp"
	r "github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	"math/rand"
	"time"
)

var (
	Dbm *gorp.DbMap
)

func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	Dbm.AddTable(models.Game{}).SetKeys(true, "Id")
	Dbm.AddTable(models.Organization{}).SetKeys(true, "Id")
	Dbm.AddTable(models.Player{}).SetKeys(true, "Id")
	Dbm.AddTable(models.User{}).SetKeys(true, "Id")
	Dbm.TraceOn("\x1b[36m[gorp]\x1b[0m", r.INFO)
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
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
