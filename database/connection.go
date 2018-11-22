package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alanwgt/apateapi/models"
	"github.com/alanwgt/apateapi/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var conn *gorm.DB

func init() {
	connect()
}

// connects to the database, it will panic if we don't have a connection
func connect() {
	dbConf := util.Conf.DB
	c, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbConf.Host, dbConf.Port, dbConf.User, dbConf.DbName, dbConf.Password))

	if err != nil {
		log.Fatal(err)
	}

	// Disable table name's pluralization globally
	c.SingularTable(true) // if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected

	c.DB().SetMaxIdleConns(util.Conf.DB.MaxIdleConns)
	c.DB().SetMaxOpenConns(util.Conf.DB.MaxOpenConns)
	c.LogMode(true)

	f, err := os.OpenFile(util.Conf.DB.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Printf("Couldn't open log file '%s'\n", util.Conf.Server.LogFile)
		log.Fatalln(err)
	}

	// sets the default logger for the database stuff
	c.SetLogger(log.New(f, "\r\n", 0666))
	// we need to manually set the postgres schema
	c.Exec("SET search_path TO " + dbConf.Schema)

	// AutoMigrate will automatically create the columns
	c.AutoMigrate(
		&models.User{},
		&models.FriendRequest{},
		&models.Blocked{},
		&models.LoginAttempt{},
		&models.Message{},
		&models.MessageContent{},
	)

	c.Model(&models.FriendRequest{}).AddForeignKey("user_id", "apate.user(id)", "CASCADE", "CASCADE")
	c.Model(&models.FriendRequest{}).AddForeignKey("request_to", "apate.user(id)", "CASCADE", "CASCADE")

	c.Model(&models.Blocked{}).AddForeignKey("user_id", "apate.user(id)", "CASCADE", "CASCADE")
	c.Model(&models.Blocked{}).AddForeignKey("blocked_id", "apate.user(id)", "CASCADE", "CASCADE")

	c.Model(&models.Message{}).AddForeignKey("user_id", "apate.user(id)", "CASCADE", "CASCADE")
	c.Model(&models.Message{}).AddForeignKey("recipient_id", "apate.user(id)", "CASCADE", "CASCADE")

	c.Model(&models.MessageContent{}).AddForeignKey("message_id", "apate.message(id)", "CASCADE", "CASCADE")

	c.Model(&models.LoginAttempt{}).AddForeignKey("user_id", "apate.user(id)", "CASCADE", "CASCADE")

	conn = c
}

// Create inserts an object into database
func Create(value interface{}) {
	log.Println("Inserting:", value)
	conn.Create(value)
}

// GetOpenConnection returns an instance of the current open connection
func GetOpenConnection() *gorm.DB {
	return conn
}
