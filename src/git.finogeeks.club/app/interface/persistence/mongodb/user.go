package mongodb

import (
	"git.finogeeks.club/app/domain/model"
	"git.finogeeks.club/app/domain/repository"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

type userRepository struct {
	colName string
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		colName: viper.GetString("mongo.collection.user"),
	}
}

func (r *userRepository) FindAll() (obj []*model.User, err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(r.colName)
	err = col.Find(bson.M{}).All(&obj)
	return
}

func (r *userRepository) FindByEmail(email string) (obj *model.User, err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(r.colName)
	err = col.Find(bson.M{"email": email}).One(&obj)
	return
}

func (r *userRepository) Save(user *model.User) (err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(r.colName)
	err = col.Insert(user)
	return
}
