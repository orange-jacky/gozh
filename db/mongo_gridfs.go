package db

import "github.com/globalsign/mgo"

type MongoGridfs struct {
	MongoClient
	GridFS *mgo.GridFS
}

func (m *MongoGridfs) C() {
	fs := "fs"
	if m.Collection != "" {
		fs = m.Collection
	}
	m.GridFS = m.Session.Db.GridFS(fs)
}

func (m *MongoGridfs) GetGridFs() *mgo.GridFS {
	return m.GridFS
}
