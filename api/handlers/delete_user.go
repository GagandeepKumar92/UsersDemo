package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

// func NewFindUser(rt *runtime) users.FindUsersHandler{
func NewDeleteUser(rt *gserver.Runtime) users.DeleteUserHandler {
	return &deleteUser{rt: rt}
}

type deleteUser struct {
	rt *gserver.Runtime
}

func (f *deleteUser) Handle(fup users.DeleteUserParams) middleware.Responder {

	//n := "Gagandeep Kumar"
	//us := []*models.User{{Address: "ABC", ID: 2, Name: &n}}
	if err := f.rt.GetManager().DeleteUser(fup.ID); err != nil {
		return users.NewDeleteUserNotFound().WithPayload(asErrorResponse(err))
	}

	//res := users.NewFindUsersOK().WithPayload(us)
	//fmt.Println(f.rt.AppName, "Find User Request Hit")

	return users.NewDeleteUserNoContent()
	// print the appication name

}
