package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

func NewAddNewUser(rt *gserver.Runtime) users.AddUserHandler {
	return &addUser{rt: rt}
}

type addUser struct {
	rt *gserver.Runtime
}

func (f *addUser) Handle(fup users.AddUserParams) middleware.Responder {

	//fup.B
	//	n := "Gagandeep Kumar (Add User example)"
	//us := []*models.User{{Address: "ABC (Add User example)", ID: 3, Name: &n}}
	//res := users.NewAddUserCreated().WithPayload(us[0])
	//fmt.Println(f.rt.AppName, "Add User Request Hit")
	usr := &domain.User{ID: fup.Body.ID, Name: *fup.Body.Name, Address: fup.Body.Address}

	if err := f.rt.GetManager().CreateUser(usr); err != nil {
		return users.NewAddUserBadRequest().WithPayload(asErrorResponse(err))
	}

	/*usResponse := []*models.User{}
	usResponse = append(usResponse, asUserResponse(usr))
	users.NewAddUserCreated().WithPayload(usResponse[0])*/

	return users.NewAddUserCreated()
	// print the appication name
}
