package db

import mgo "gopkg.in/mgo.v2"

type DBObject struct {
}

func NewO() DBObject {
	return DBObject{}
}

func (d DBObject) NewDB() *mgo.Session {
	return &mgo.Session{}
}

func (d DBObject) Get() string {
	return "normal db"
}
