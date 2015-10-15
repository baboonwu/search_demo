package v2

import (
	"github.com/ant0ine/go-json-rest/rest"
	"tapi_go/method"
	"tapi_go/model/v2"
	"tapi_go/route/middleware"
	"tapi_go/tools"
)

func SearchRoute() []*rest.Route {
	return []*rest.Route{
		// search
		middleware.MakeAuthRouter(method.GET, tools.VersionV2("/tsearch/users"), SearchUsers, false),
	}
}

func SearchUsers(w rest.ResponseWriter, r *rest.Request) {

	type RespSearchUser struct {
		User model.UserNet `json:"user"`
	}

	resp := RespSearchUser{User: model.UserNet{Id: "123"}}

	w.WriteJson(tools.SuccessJson(resp))
	return
}
