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
	dbm.TraceOn("[gorp]", r.INFO)

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

	/*
			setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
				for col, size := range colSizes {
					t.ColMap(col).MaxSize = size
				}
			}

			t := dbm.AddTable(models.User{}).SetKeys(true, "UserId")
			t.ColMap("Password").Transient = true
			setColumnSizes(t, map[string]int{
				"Username": 20,
				"Name":     100,
			})

			t = dbm.AddTable(models.Hotel{}).SetKeys(true, "HotelId")
			setColumnSizes(t, map[string]int{
				"Name":    50,
				"Address": 100,
				"City":    40,
				"State":   6,
				"Zip":     6,
				"Country": 40,
			})

			t = dbm.AddTable(models.Booking{}).SetKeys(true, "BookingId")
			t.ColMap("User").Transient = true
			t.ColMap("Hotel").Transient = true
			t.ColMap("CheckInDate").Transient = true
			t.ColMap("CheckOutDate").Transient = true
			setColumnSizes(t, map[string]int{
				"CardNumber": 16,
				"NameOnCard": 50,
			})

		dbm.TraceOn("[gorp]", r.INFO)
		dbm.CreateTables()

		bcryptPassword, _ := bcrypt.GenerateFromPassword(
			[]byte("demo"), bcrypt.DefaultCost)
			demoUser := &models.User{0, "Demo User", "demo", "demo", bcryptPassword}
			if err := dbm.Insert(demoUser); err != nil {
				panic(err)
			}

			hotels := []*models.Hotel{
				&models.Hotel{0, "Marriott Courtyard", "Tower Pl, Buckhead", "Atlanta", "GA", "30305", "USA", 120},
				&models.Hotel{0, "W Hotel", "Union Square, Manhattan", "New York", "NY", "10011", "USA", 450},
				&models.Hotel{0, "Hotel Rouge", "1315 16th St NW", "Washington", "DC", "20036", "USA", 250},
			}
			for _, hotel := range hotels {
				if err := dbm.Insert(hotel); err != nil {
					panic(err)
				}
			}
	*/
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
