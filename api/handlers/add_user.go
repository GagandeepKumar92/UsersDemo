package handlers

import (
	gserver "GaganSimpleServer"
	"GaganSimpleServer/domain"
	"GaganSimpleServer/gen/restapi/operations/users"

	"github.com/go-openapi/runtime/middleware"
)

func NewAddNewUser(rt *gserver.Runtime) users.AddUserHandler {
	return &addUser{rt: rt}
}

type addUser struct {
	rt *gserver.Runtime
}

func (f *addUser) Handle(fup users.AddUserParams) middleware.Responder {

	usr := &domain.User{ID: fup.Body.ID, Name: *fup.Body.Name, Address: fup.Body.Address}

	if err := f.rt.GetManager().CreateUser(usr); err != nil {
		return users.NewAddUserBadRequest().WithPayload(asErrorResponse(err))
	}

	return users.NewAddUserCreated()

}
