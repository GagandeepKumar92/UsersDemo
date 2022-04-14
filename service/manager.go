package service

import (
	"fmt"
	"time"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/segmentio/ksuid"
)

// Add User

type mgr struct {
	ds db.DataStore
	//ds map[int]*domain.User
}

func (m *mgr) CreateUser(usr *domain.User) *domain.Error {

	//usr.CreatedAt = time.Now().UTC()
	usr.CreatedAt = time.Now().UTC()
	usr.ID = ksuid.New().String()

	if len(usr.Name) < 3 {
		return &domain.Error{Code: 400, Message: "Name should be at least 3 characters long"}
	}

	m.ds.AddUser(usr)
	return nil

	//m.ds[usr.ID] = usr
	// store data in (m.ds)
}

func (m *mgr) ViewUser(i string) (*domain.User, error) {
	//fmt.Printf("The type of id is %T",fmt.Sprint(i))
	// user,err := m.ds.ViewUser(i)
	// if err != nil{
	// 	return nil,&domain.Error{Code:404,Message:"User doesn't exist"}
	// }
	return m.ds.ViewUser(i)
}

func (m *mgr) UpdateUser(usr *domain.User) error {

	fmt.Println(usr.ID)
	//usr.CreatedAt = time.Now().UTC()
	dbUser, ViewErr := m.ds.ViewUser(usr.ID)

	if ViewErr != nil {
		fmt.Println("1")
		return &domain.Error{Code: 404, Message: "User doesn't exist"}
	}

	dbUser.Address = usr.Address
	/*if err := m.ds.UpdateUser(usr); err != nil {
		return &domain.Error{Code: 404, Message: "User doesn't exist"}
	}*/
	return m.ds.UpdateUser(dbUser)

	//m.ds[usr.ID] = usr
	// store data in (m.ds)
}

func (m *mgr) DeleteUser(id string) *domain.Error {

	if err := m.ds.DeleteUser(id); err != nil {
		return &domain.Error{Code: 404, Message: "User doesn't exist"}
	}
	return nil
}

func (m *mgr) ListUser(limit int32, name string) []*domain.User {
	// iterate (m.ds) and send all users from here

	user, err := m.ds.ListUsers(limit, name)
	if err != nil {
		return nil
	}
	return user

	/*s := make([]*domain.User, 0)
	for _, element := range m.ds {
		s = append(s, element)
		//fmt.Println(s1)
		//	fmt.Println("Key:", key, "=>", "Element:", element)
	}*/
	/*for i := 0; i < len(m.ds); i++ {
		users := m.ds[usr.ID]
	}*/

	//	res := models.users.NewFindUsersOK().WithPayload(us)
	//return s
	//ret
}
