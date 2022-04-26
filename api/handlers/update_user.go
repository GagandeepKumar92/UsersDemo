package handlers

import (
	"fmt"

	gserver "GaganSimpleServer"
	"GaganSimpleServer/domain"
	"GaganSimpleServer/gen/restapi/operations/users"

	"github.com/go-openapi/runtime/middleware"
)

func NewUpdateUser(rt *gserver.Runtime) users.UpdateUserHandler {
	return &updateUser{rt: rt}
}

type updateUser struct {
	rt *gserver.Runtime
}

func (f *updateUser) Handle(fup users.UpdateUserParams) middleware.Responder {

	fmt.Println("User Id is ", fup.ID)

	if *fup.Body.Address == "" {
		return users.NewUpdateUserDefault(422)
	}

	usr := &domain.User{ID: fup.ID, Address: *fup.Body.Address}

	if err := f.rt.GetManager().UpdateUser(usr); err != nil {

		derr, ok := err.(domain.Err)

		if ok {
			switch derr.StatusCode() {
			case 404:
				return users.NewUpdateUserNotFound().WithPayload(asErrorResponse(err.(*domain.Error)))
			}
		} else {
			return users.NewUpdateUserDefault(500).WithPayload(asErrorResponse(&domain.Error{Message: "Internal Server Error"}))
		}

	}

	return users.NewUpdateUserNoContent()
}
