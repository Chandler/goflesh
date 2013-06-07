package controllers

import (
	"database/sql"
	"flesh/app/models"
	"fmt"
	"github.com/coopernurse/gorp"
	r "github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	"math/rand"
	"time"
)

var (
	dbm *gorp.DbMap
)

func Init() {
	db.Init()
	dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	dbm.AddTable(models.Organization{}).SetKeys(true, "Id")
	dbm.AddTable(models.Game{}).SetKeys(true, "Id")
	dbm.AddTable(models.User{}).SetKeys(true, "Id")
	dbm.AddTable(models.Player{}).SetKeys(true, "Id")
	dbm.TraceOn("\x1b[36m[gorp]\x1b[0m", r.INFO)
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
