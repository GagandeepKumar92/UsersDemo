package inmemory

import (
	"fmt"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
)

func init() {
	db.RegisterDataStore("inmemory", NewClient)
	fmt.Println("Call is in Init")
}

func NewClient() (db.DataStore, error) {
	return &client{ds: make(map[string]*domain.User)}, nil
}

type client struct {
	ds map[string]*domain.User
}

func (c *client) ViewUser(id string) (*domain.User, error) {

	var userInfo *domain.User
	var ok bool
	if userInfo, ok = c.ds[id]; !ok {
		return nil, &domain.Error{Code: 404, Message: "User doesn't exist"}
	}

	return userInfo, nil
}

func (c *client) UpdateUser(user *domain.User) error {

	//idfound := false

	c.ds[user.ID] = user

	/*if !idfound {
		return &domain.Error{Code: 404, Message: "User doesn't exist"}
	}*/

	return nil
}

func (c *client) AddUser(user *domain.User) (string, error) {
	c.ds[user.ID] = user
	return user.ID, nil
}

func (c *client) ListUsers(limit int32, name string) ([]*domain.User, error) {

	var userInfo = []*domain.User{}
	for _, value := range c.ds {
		userInfo = append(userInfo, value)
	}
	fmt.Println("The value of userInfo is", userInfo)
	return userInfo, nil
}

func (c *client) DeleteUser(id string) error {
	if _, ok := c.ds[id]; ok {
		delete(c.ds, id)
		return nil
	}
	return nil
}
