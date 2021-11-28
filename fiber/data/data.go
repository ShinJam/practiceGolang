package data

import (
	"fmt"

	"xorm.io/xorm"
)

type connectionInfo struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

func (db connectionInfo) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", db.host, db.port, db.user, db.password, db.dbname, db.sslmode)
}

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string `json:"-"`
}

func CreateDBEngine() (*xorm.Engine, error) {
	connectionInfo := connectionInfo{host: "localhost", port: 5432, user: "root", password: "password", dbname: "authServer", sslmode: "disable"}

	engine, err := xorm.NewEngine("postgres", connectionInfo.String())
	if err != nil {
		return nil, err
	}
	if err := engine.Ping(); err != nil {
		return nil, err
	}

	if err := engine.Sync(new(User)); err != nil {
		return nil, err
	}

	return engine, nil
}
