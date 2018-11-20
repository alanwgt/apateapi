package routes

import (
	"net/http"

	"github.com/alanwgt/apateapi/controllers"
	"github.com/gorilla/mux"
)

// Route type holds all the information needed about any route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of Route
type Routes []Route

// BuildRouter creates a *mux.Router and associate all the registered routes to it
func BuildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"CreateAccount",
		"POST",
		"/user",
		controllers.CreateAccount,
	},
	Route{
		"UserHandshake",
		"POST",
		"/user/handshake",
		controllers.Handshake,
	},
	Route{
		"ServerPublickey",
		"GET",
		"/server/pubk",
		controllers.GetServerPubK,
	},
	Route{
		"QueryUser",
		"GET",
		"/user/q/{username}",
		controllers.QueryUsers,
	},
	Route{
		"DeleteAccount",
		"DELETE",
		"/user",
		controllers.DeleteAccount,
	},
}
