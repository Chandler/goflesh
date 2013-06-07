package tests

import (
	"database/sql"
	"flesh/app/models"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	"io/ioutil"
)

var (
	dbm         *gorp.DbMap
	create_conn *sql.DB
	db_spec     string
	driver      string
	drop_query  string
	copy_query  string
	testMode    bool
	db_exists   bool
)

func TestInit() {
	defer MakeDbFromTemplate()

	if dbm != nil {
		return
	}

	// Read configuration.
	var found bool
	driver, found = revel.Config.String("db.driver")
	if !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	db_spec, found = revel.Config.String("db.spec")
	if !found {
		revel.ERROR.Fatal("No db.spec found.")
	}
	create_spec, found := revel.Config.String("db.template.create_spec")
	if !found {
		revel.ERROR.Fatal("No db.template.create_spec found.")
	}
	template_spec, found := revel.Config.String("db.template.spec")
	if !found {
		revel.ERROR.Fatal("No db.template.spec found.")
	}
	template_db_name, found := revel.Config.String("db.template.name")
	if !found {
		revel.ERROR.Fatal("No db.template.name found.")
	}
	test_db_name, found := revel.Config.String("db.template.to_name")
	if !found {
		revel.ERROR.Fatal("No db.template.to_name found.")
	}

	drop_query = fmt.Sprintf(`
	DROP DATABASE %s;
	`, test_db_name)

	copy_query = fmt.Sprintf(`
	CREATE DATABASE %s TEMPLATE %s
	`, test_db_name, template_db_name)

	// Slurp in the DB schema
	sql_file_bytes, err := ioutil.ReadFile("tests/testdb.sql")
	if err != nil {
		revel.ERROR.Fatal(err)
	}
	sql_file_string := string(sql_file_bytes)

	// Open a connection to drop and create the DB
	create_conn, err = sql.Open(driver, create_spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	// Drop the template database, if there is one
	drop_template_query := "DROP DATABASE " + template_db_name
	_, err = create_conn.Exec(drop_template_query)
	if err != nil {
		revel.INFO.Print(err)
	}

	// Drop the test database, if there is one
	db.Db.Close()
	_, err = create_conn.Exec(drop_query)
	if err != nil {
		revel.INFO.Print(err)
	}

	// Create the template database
	create_template_query := "CREATE DATABASE " + template_db_name
	_, err = create_conn.Exec(create_template_query)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	// Open a connection to the template DB
	template_conn, err := sql.Open(driver, template_spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	// Build the template database from SQL dump
	_, err = template_conn.Exec(sql_file_string)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	// Close template connection
	err = template_conn.Close()
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	testMode, _ := revel.Config.Bool("mode.test")
	_ = testMode // testMode used by other testing modules
}

func MakeDbFromTemplate() {
	var err error

	if db_exists {
		TestClean()
	}

	_, err = create_conn.Exec(copy_query)
	if err != nil {
		revel.ERROR.Fatal(err)
	}
	db_exists = true

	db.Db, err = sql.Open(driver, db_spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}
	dbm.AddTable(models.Organization{}).SetKeys(true, "Id")
	dbm.AddTable(models.Game{}).SetKeys(true, "Id")
	dbm.AddTable(models.User{}).SetKeys(true, "Id")
	dbm.AddTable(models.Player{}).SetKeys(true, "Id")
	dbm.TraceOn("\x1b[36m[gorp]\x1b[0m", revel.INFO)

}

func TestClean() {
	err := db.Db.Close()
	if err != nil {
		revel.ERROR.Print(err)
	}

	_, err = create_conn.Exec(drop_query)
	if err != nil {
		revel.ERROR.Fatal(err)
	}
	db_exists = false

}
