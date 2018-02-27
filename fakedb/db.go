package db

import mgo "gopkg.in/mgo.v2"

type FakeDBObject struct {
}

func NewO() FakeDBObject {
	return FakeDBObject{}
}

func (d FakeDBObject) NewDB() *mgo.Session {
	return &mgo.Session{}
}

func (d FakeDBObject) Get() string {
	return "fake db"
}
