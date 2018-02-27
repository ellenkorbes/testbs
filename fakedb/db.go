package db

import mgo "gopkg.in/mgo.v2"

type FakeDBObject struct {
}

func NewDBObject() FakeDBObject {
	return FakeDBObject{}
}

func (d FakeDBObject) NewSession() *mgo.Session {
	return &mgo.Session{}
}

func (d FakeDBObject) Get() string {
	return "fake db"
}
