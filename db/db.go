package db

import mgo "gopkg.in/mgo.v2"

type DBObject struct {
}

func NewDBObject() DBObject {
	return DBObject{}
}

func (d DBObject) NewSession() *mgo.Session {
	return &mgo.Session{}
}

func (d DBObject) Get() string {
	return "normal db"
}
