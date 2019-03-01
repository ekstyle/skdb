package routes

import (
	"fmt"
	"github.com/ekstyle/skdb/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"strings"
)

var routes = Routes{

	Route{
		"Main",
		"GET",
		"",
		"",
		"/api",
		controllers.YourHandler,
	},
}

func HandlerFs(publicDir string) http.Handler {
	handler := http.FileServer(http.Dir(publicDir))
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("here!")
		_path := req.URL.Path
		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, req)
			return
		}
		// the all 404 gonna be served as root
		http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		if route.Queries == "" {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler).
				Queries(route.Queries, route.Hadle)
		}
	}
	//Serve static content
	router.PathPrefix("/").Handler(HandlerFs("./public"))
	return router
}
