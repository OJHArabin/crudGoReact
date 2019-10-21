package databaseConfig

import mgo "gopkg.in/mgo.v2"

func DbDetails() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func CrudCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("crud_db").C("crud_c")
}
