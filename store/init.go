package store

import (
	"github.com/ekstyle/skdb/params"
	"gopkg.in/mgo.v2"
	"log"
)

func getDB() func() *mgo.Session {
	var sesion *mgo.Session
	return func() *mgo.Session {
		if sesion == nil {
			var err error
			sesion, err = mgo.Dial(params.MongoUrl())
			if err != nil {
				log.Fatal("MONGO error: " + err.Error())
			}
			return sesion.Clone()
		}
		return sesion.Clone()
	}
}

func init() {
	db := getDB()
	defer db().Close()
}
