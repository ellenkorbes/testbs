package main

import (
	"github.com/ellenkorbes/testbs/ctrl"
	// "github.com/ellenkorbes/testbs/db"
	db "github.com/ellenkorbes/testbs/fakedb"
)

func main() {
	d := db.NewSession()
	c := ctrl.NewController(d)

	c.Get()

}
