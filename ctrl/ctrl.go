package ctrl

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

type DBI interface {
	NewDB() *mgo.Session
	Get() string
}

type Controller struct {
	DB    DBI
	Mongo *mgo.Session
}

func NewController(db DBI) *Controller {
	return &Controller{
		DB:    db,
		Mongo: db.NewDB(),
	}
}

func (c *Controller) Get() {
	x := c.DB.Get()
	fmt.Println(x)
}
