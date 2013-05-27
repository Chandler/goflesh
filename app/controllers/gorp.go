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

type GorpPlugin struct {
	r.EmptyPlugin
}

func (p GorpPlugin) OnAppStart() {
	db.DbPlugin{}.OnAppStart()
	dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	dbm.AddTable(models.Organization{}).SetKeys(true, "Id")
	dbm.AddTable(models.Game{}).SetKeys(true, "Id")
	dbm.AddTable(models.User{}).SetKeys(true, "Id")
	dbm.AddTable(models.Player{}).SetKeys(true, "Id")
	dbm.TraceOn("\x1b[36m[gorp]\x1b[0m", r.INFO)

	// Create tables (ok if they exist, move on)
	// TODO: replace this with official schema
	dbm.CreateTables()

	organizations := []*models.Organization{
		&models.Organization{0, "UIdaho", "vandals", "US/Pacific", nil, nil},
		&models.Organization{0, "Boise State", "broncs", "US/Mountain", nil, nil},
		&models.Organization{0, "Berkeley", "cal", "US/Pacific", nil, nil},
	}

	for _, organization := range organizations {
		if err := dbm.Insert(organization); err != nil {
			// panic(err)
		}
		randGameNum := rand.Int()
		now := time.Now().UTC()
		later := now.Add(12 * time.Hour)
		tomorrow := now.Add(24 * time.Hour)
		tomorrowLater := later.Add(24 * time.Hour)
		game := &models.Game{0,
			fmt.Sprintf("Game number %d", randGameNum),
			fmt.Sprintf("game-%d", randGameNum),
			organization.Id,
			organization.Default_timezone,
			&now,
			&tomorrow, //&time.Now() + time.Hour,
			&later,
			&tomorrowLater,
			nil,
			nil,
		}
		if err := dbm.Insert(game); err != nil {
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
