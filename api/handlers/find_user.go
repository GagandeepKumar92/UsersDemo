package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/models"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

// func NewFindUser(rt *runtime) users.FindUsersHandler{
func NewFindUser(rt *gserver.Runtime) users.FindUsersHandler {
	return &findUser{rt: rt}
}

type findUser struct {
	rt *gserver.Runtime
}

func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder {
	//if fup.Name != nil {
	fmt.Println("Name = ", fup.Name, "Limit = ", fup.Limit)
	//}
	//n := "Gagandeep Kumar"
	//us := []*models.User{{Address: "ABC", ID: 2, Name: &n}}
	var limit int32
	var name string
	var us []*domain.User
	if fup.Limit != nil && fup.Name != nil {
		limit = *fup.Limit
		name = *fup.Name
		us = f.rt.GetManager().ListUser(limit, name)
	} else if fup.Limit != nil {
		limit = *fup.Limit
		us = f.rt.GetManager().ListUser(limit, "")
	} else if fup.Name != nil {
		name = *fup.Name
		us = f.rt.GetManager().ListUser(0, name)
	} else {
		us = f.rt.GetManager().ListUser(0, "")
	}

	/*if fup.Name != nil {
		name = *fup.Name
		fmt.Println("Query Parameteres Name =", *fup.Name)
	}*/

	//us := f.rt.GetManager().ListUser(limit, name)

	usResponse := []*models.User{}
	for _, usr := range us {
		usResponse = append(usResponse, asUserResponse(usr))
	}
	res := users.NewFindUsersOK().WithPayload(usResponse)
	//fmt.Println(f.rt.AppName, "Find User Request Hit")
	return res
	// print the appication name
}
