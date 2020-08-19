package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// device
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Db - defines a sesion of the database
var Db *gorm.DB

// DBConfigParams - contains the user configuration to start the database
type DBConfigParams struct {
	user    string
	pass    string
	dbname  string
	host    string
	options string
}

// newDefs - initialize a new defaultParams
func (d *defaultParams) newDefs() {
	d.user = "root"
	d.pass = ""
	d.dbname = "systemordersystem"
	d.host = "localhost:3306"
	d.options = "charset=utf8&parseTime=True&loc=Local"
}

// InitDB - creates a season to the database
func InitDB(cnf DBConfigParams) error {
	paramsParser(&cnf)
	connection := fmt.Sprintf("%s:%s@(%s)/%s?%s", cnf.user, cnf.pass, cnf.host, cnf.dbname, cnf.options)

	db, err := gorm.Open("mysql", connection)
	if err != nil {
		log.Println("Failed to connect database")
		log.Panic(err)
		return err
	}
	log.Println("--- Database Connected ---")
	Db = db
	Db.LogMode(true)
	return nil
}

// MigrateAll - execute a migration in all schemas implemented
// in the model package
func MigrateAll() {
	log.Println("--- Migrating Database ---")
	if Db == nil {
		log.Panicln("Database is null")
	} else if Db.DB().Ping() != nil {
		log.Panicln("Database sesion is down")
	} else {
		Db.AutoMigrate(&SchemaUser{})
		Db.AutoMigrate(&SchemaTechnician{})
		Db.AutoMigrate(&SchemaTelevision{})
		Db.Model(&SchemaOrder{}).AddForeignKey(
			"id",
			"schema_users(id)",
			"CASCADE",
			"CASCADE",
		)
		Db.Model(&SchemaOrder{}).AddForeignKey(
			"id",
			"schema_technicians(id)",
			"CASCADE",
			"CASCADE",
		)
		Db.Model(&SchemaOrder{}).AddForeignKey(
			"id",
			"schema_televisions(id)",
			"CASCADE",
			"CASCADE",
		)
		Db.AutoMigrate(&SchemaOrder{})
		log.Println("--- Migration Completed ---")
	}
}

// CloseDB - close the conection to the database
func CloseDB() {
	if Db != nil || Db.DB().Ping() != nil {
		Db.Close()
	}
}

// defaultParams - contains the default configuration to start the database
type defaultParams struct {
	user    string
	pass    string
	dbname  string
	host    string
	options string
}

// paramsParser - parse empty parameters
func paramsParser(cnf *DBConfigParams) {
	defs := defaultParams{}
	defs.newDefs()
	if cnf == nil {
		*cnf = DBConfigParams{}
	}
	assigner(&cnf.user, defs.user)
	assigner(&cnf.pass, defs.pass)
	assigner(&cnf.host, defs.host)
	assigner(&cnf.dbname, defs.dbname)
	assigner(&cnf.options, defs.options)

}

// assigner - assings a default value to the input if it is empty
func assigner(input *string, def string) {
	if *input == "" {
		*input = def
	}
}
