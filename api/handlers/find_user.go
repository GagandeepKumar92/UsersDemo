package handlers

import (
	"fmt"

	"GaganSimpleServer/domain"
	"GaganSimpleServer/gen/models"
	"GaganSimpleServer/gen/restapi/operations/users"

	gserver "GaganSimpleServer"

	"github.com/go-openapi/runtime/middleware"
)

func NewFindUser(rt *gserver.Runtime) users.FindUsersHandler {
	return &findUser{rt: rt}
}

type findUser struct {
	rt *gserver.Runtime
}

func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder {

	us, err := f.rt.GetManager().ListUser(*fup.Limit, filteredMap(fup))

	if err != nil {
		fmt.Println(err)
		return users.NewFindUsersDefault(500).WithPayload(asErrorResponse(&domain.Error{Message: "Internal Server Error"}))
	}

	usResponse := []*models.User{}
	for _, usr := range us {
		usResponse = append(usResponse, asUserResponse(usr))
	}

	return users.NewFindUsersOK().WithPayload(usResponse)
}

func filteredMap(fup users.FindUsersParams) map[string]interface{} {
	filterMap := make(map[string]interface{})

	if fup.Name != nil {
		filterMap["name"] = *fup.Name
	}

	return filterMap
}
