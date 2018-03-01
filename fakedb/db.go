package db

import mgo "gopkg.in/mgo.v2"

type FakeDBObject struct {
	Session *mgo.Session
}

func NewSession() FakeDBObject {
	return FakeDBObject{&mgo.Session{}}
}

func (d FakeDBObject) Get() string {
	return "fake db"
}
