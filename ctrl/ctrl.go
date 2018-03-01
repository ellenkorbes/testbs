package ctrl

import (
	"fmt"
)

type DBI interface {
	// NewSession() *mgo.Session
	Get() string
}

type Controller struct {
	DB DBI
	// Mongo *mgo.Session
}

func NewController(db DBI) *Controller {
	return &Controller{
		DB: db,
		// Mongo: db.NewSession(),
	}
}

func (c *Controller) Get() {
	x := c.DB.Get()
	fmt.Println(x)
}
