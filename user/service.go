package user

import (
	"github.com/wongoo/go-app-example/mongo"
)

func AddUser(user *User) error {
	return mongo.InsertCollection(USER_COLLECTION, user)
}

func FindUser(id string) (user *User, err error) {
	user = &User{}
	err = mongo.FindOne(USER_COLLECTION, id, user)
	if err != nil {
		return nil, err
	}
	return user, err;
}

func DeleteUser(id string) error {
	return mongo.DeleteOne(USER_COLLECTION, id)
}
