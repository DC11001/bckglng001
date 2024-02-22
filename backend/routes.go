package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Defining the type route and routes as an array of route
type Route struct {
	Name       string
	MethodType string
	Path       string
	Handler    http.HandlerFunc
}

type Routes []Route

//This func return the mux.Router used in the http.ListenAndServe
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//This loop trough the array setting the routes to the mux router
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.MethodType).
			Path(route.Path).
			Handler(route.Handler)
	}
	return router
}

//Examples routes with basic http methods. The functions are called from the actions file.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"CreateUser",
		"POST",
		"/createUser",
		CreateUserHandler,
	},
	Route{
		"ModifyUser",
		"POST",
		"/modifyUser",
		ModifyUserHandler,
	},
	Route{
		"CreateRoom",
		"POST",
		"/createRoom",
		CreateRoomHandler,
	},
	Route{
		"ModifyRoom",
		"POST",
		"/modifyRoom",
		ModifyRoomHandler,
	},
	Route{
		"DeleteRoom",
		"DELETE",
		"/deleteRoom",
		DeleteRoomHandler,
	},
	Route{
		"SearchRoom",
		"POST",
		"/searchRoom",
		SearchRoomHandler,
	},
	Route{
		"SearchRoomByState",
		"POST",
		"/searchRoomByState",
		SearchRoomByStateHandler,
	},
	Route{
		"CreateReservation",
		"POST",
		"/createReservation",
		CreateReservationHandler,
	},
	Route{
		"ModifyReservation",
		"POST",
		"/modifyReservation",
		ModifyReservationHandler,
	},
	Route{
		"DeleteReservation",
		"DELETE",
		"/deleteReservation",
		DeleteReservationHandler,
	},
	Route{
		"ClientReport",
		"GET",
		"/clientReport",
		ClientReportHandler,
	},
	Route{
		"DeleteRoute",
		"DELETE",
		"/delete",
		Delete,
	},
}
