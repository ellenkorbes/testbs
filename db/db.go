package db

import mgo "gopkg.in/mgo.v2"

type DBObject struct {
	Session *mgo.Session
}

func NewSession() DBObject {
	return DBObject{&mgo.Session{}}
}

func (d DBObject) Get() string {
	return "normal db"
}
