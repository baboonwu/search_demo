package route

import (
	"github.com/ant0ine/go-json-rest/rest"
	"tsearch/route/v2"
)

func GetApiSearch() []*rest.Route {

	routes := []*rest.Route{}

	// addSearchRoutes(&routes, healthcheck())
	addSearchRoutes(&routes, v2.SearchRoute())

	return routes
}

// 增加各个模块的route映射关系
func addSearchRoutes(routes *[]*rest.Route, routeArray []*rest.Route) {
	for _, route := range routeArray {
		*routes = append(*routes, route)
	}
}
