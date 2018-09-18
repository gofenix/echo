package mongodb

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

type (
	Session struct {
		*mgo.Session
	}

	Collection struct {
		*mgo.Collection
	}
)

var dbses Session

func init() {

	if viper.GetString("mongo.domain") != "" {
		InitMongo()
	}
}

func InitMongo() (err error) {
	domain := viper.GetString("mongo.domain")
	port := viper.GetString("mongo.port")
	db := viper.GetString("mongo.db")
	auth := viper.GetString("mongo.auth")
	url := "mongodb://" + auth + "@" + domain + port + "/" + db
	fmt.Println("url: ", url)
	dbses.Session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	dbses.SetPoolLimit(50)
	return nil
}

func EnsureIndex(col Collection, primaryKeys []string, unique bool) (err error) {
	index := mgo.Index{
		Key:        primaryKeys,
		Unique:     unique,
		DropDups:   unique,
		Background: true,
		Sparse:     true,
	}
	err = col.EnsureIndex(index)
	return
}

func (ses Session) GetDBCol(colName string) (col Collection) {
	col.Collection = ses.DB(viper.GetString("mongo.db")).C(colName)
	return
}

func (ses Session) UninitDB() {
	ses.Close()
}
