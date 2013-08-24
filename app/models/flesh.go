package models

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	"time"
)

var (
	Dbm *gorp.DbMap
)

func Init() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}
	Dbm.TraceOn("\x1b[36m[M.gorp]\x1b[0m", revel.INFO)
	AddTables(Dbm)
}

func AddTables(dbm *gorp.DbMap) {
	dbm.AddTable(Game{}).SetKeys(true, "Id")
	dbm.AddTable(Organization{}).SetKeys(true, "Id")
	dbm.AddTable(Player{}).SetKeys(true, "Id")
	dbm.AddTable(User{}).SetKeys(true, "Id")
	dbm.AddTable(Oz{}).SetKeys(true, "Id")
	dbm.AddTable(Tag{}).SetKeys(true, "Id")
	dbm.AddTable(Member{}).SetKeys(true, "Id")
	dbm.AddTable(Event{}).SetKeys(true, "Id")
	dbm.AddTable(EventType{}).SetKeys(true, "Id")
	dbm.AddTable(EventRole{}).SetKeys(true, "Id")
	dbm.AddTable(EventToPlayer{}).SetKeys(true, "Id")
	dbm.AddTableWithName(HumanCode{}, "human_code").SetKeys(false, "Id")
	dbm.AddTableWithName(OzPool{}, "oz_pool").SetKeys(true, "Id")
	dbm.AddTableWithName(PasswordReset{}, "password_reset")
}

type TimeTrackedModel struct {
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

func (model *TimeTrackedModel) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Created = &now
	model.Updated = model.Created
	return nil
}

func (model *TimeTrackedModel) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Updated = &now
	return nil
}
