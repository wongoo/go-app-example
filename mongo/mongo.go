package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"

	"github.com/golang/glog"
	"github.com/wongoo/go-app-example/constants"
)

var session *mgo.Session

func init() {
	session = GetMongoClient()
}

func GetMongoClient() *mgo.Session {
	if nil == session {
		dialInfo := &mgo.DialInfo{
			Addrs:    constants.MongoDbHosts,
			Database: constants.MongoDbName,
			Username: constants.MongoDbUser,
			Password: constants.MongoDbPassword,
		}
		fmt.Printf("dialInfo: %+v\n", dialInfo)

		var err error
		session, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			glog.Infof("Build MongoDB client err : %v", err.Error())
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
		session.SetPoolLimit(constants.MongoDbPoolLimit)
	}

	return session
}

func InsertDbCollection(db string, collection string, doc interface{}) error {
	fmt.Printf("insert %v.%v\n", db, collection)

	session := GetMongoClient().Clone()
	defer session.Close()

	c := session.DB(db).C(collection)
	err := c.Insert(doc)
	return err
}

func InsertCollection(collection string, doc interface{}) error {
	return InsertDbCollection(constants.MongoDbName, collection, doc)
}

func FindOne(collection string, id string, doc interface{}) (err error) {
	session := GetMongoClient().Clone()
	defer session.Close()

	c := session.DB(constants.MongoDbName).C(collection)
	query := c.FindId(id)
	err = query.One(doc)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOne(collection string, id string) error {
	session := GetMongoClient().Clone()
	defer session.Close()

	c := session.DB(constants.MongoDbName).C(collection)
	err := c.RemoveId(id)
	return err
}
