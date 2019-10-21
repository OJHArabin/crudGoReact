package model

import (
	"../databaseConfig"
	"github.com/gofrs/uuid"

	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func InsertData(deserilize []UserDetails, serilizesChannel chan []UserDetails, session *mgo.Session) ([]UserDetails, error) {
	for i := 0; i < len(deserilize); i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte(deserilize[i].Password), bcrypt.DefaultCost)
		deserilize[i].Password = string(hash)
		id, _ := uuid.NewV4()
		deserilize[i].Id = id.String()
		c := databaseConfig.CrudCollection(session)
		err := c.Insert(&deserilize[i])
		if err != nil {
			return []UserDetails{}, err
		}
	}

	return deserilize, nil
}

func DeleteUser(id string, session *mgo.Session) error {
	c := databaseConfig.CrudCollection(session)
	err := c.Remove(bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, data UserDetails, session *mgo.Session) (UserDetails, error) {
	c := databaseConfig.CrudCollection(session)

	var d bson.M
	if data.UserName != "" && data.FullName != "" && data.Email != "" {
		d = bson.M{"$set": bson.M{"username": data.UserName, "fullname": data.FullName, "email": data.Email}}
	} else if data.UserName == "" && data.FullName != "" && data.Email != "" {
		d = bson.M{"$set": bson.M{"fullname": data.FullName, "email": data.Email}}
	} else if data.UserName != "" && data.FullName == "" && data.Email != "" {
		d = bson.M{"$set": bson.M{"username": data.UserName, "email": data.Email}}
	} else if data.UserName != "" && data.FullName != "" && data.Email == "" {
		d = bson.M{"$set": bson.M{"username": data.UserName, "fullname": data.FullName}}
	} else if data.UserName == "" && data.FullName == "" && data.Email != "" {
		d = bson.M{"$set": bson.M{"email": data.Email}}
	} else if data.UserName != "" && data.FullName == "" && data.Email == "" {
		d = bson.M{"$set": bson.M{"username": data.UserName}}
	} else if data.UserName == "" && data.FullName != "" && data.Email == "" {
		d = bson.M{"$set": bson.M{"fullname": data.FullName}}
	}
	err := c.Update(bson.M{"id": id}, d)
	if err == nil {
		var ud UserDetails
		err = c.Find(bson.M{"id": id}).One(&ud)
		if err != nil {
			return UserDetails{}, err
		}
		return ud, nil
	}
	return UserDetails{}, err

}

func GetUser(id string, session *mgo.Session) (UserDetails, error) {
	var userdetails UserDetails
	c := databaseConfig.CrudCollection(session)
	err := c.Find(bson.M{"id": id}).One(&userdetails)
	if err != nil {
		return UserDetails{}, err
	}
	return userdetails, nil
}

func GetUserAll(session *mgo.Session) ([]UserDetails, error) {
	var userdetails []UserDetails
	c := databaseConfig.CrudCollection(session)
	err := c.Find(nil).All(&userdetails)
	if err != nil {
		return []UserDetails{}, err
	}
	return userdetails, nil
}

func CheckUserNameAndPassword(username, password string, session *mgo.Session) bool {
	var userdetails UserDetails
	c := databaseConfig.CrudCollection(session)
	err := c.Find(bson.M{"username": username}).One(&userdetails)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(userdetails.Password), []byte(password)) == nil
	}

}
