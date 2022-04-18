package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/models"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

func NewFindUser(rt *gserver.Runtime) users.FindUsersHandler {
	return &findUser{rt: rt}
}

type findUser struct {
	rt *gserver.Runtime
}

func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder {

	us := f.rt.GetManager().ListUser(*fup.Limit, filteredMap(fup.Name))

	usResponse := []*models.User{}
	for _, usr := range us {
		usResponse = append(usResponse, asUserResponse(usr))
	}
	res := users.NewFindUsersOK().WithPayload(usResponse)

	return res

}

func filteredMap(fileteredName *string) map[string]interface{} {
	filterMap := make(map[string]interface{})

	if fileteredName != nil {
		filterMap["name"] = *fileteredName
	}

	return filterMap
}
