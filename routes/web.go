package routes

import (
	_ "flag"

	"github.com/gorilla/mux"
	"github.com/radenaryayusuf/todoappweb/controllers"
)

func Init() *mux.Router {

	// var css, js string

	// flag.StringVar(&css, "/web/static/css", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.StringVar(&js, "/web/static/js", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()
	route := mux.NewRouter()
	// s := http.StripPrefix("/web/static/css/", http.FileServer(http.Dir("./web/static/css/")))
	// // This will serve files under http://localhost:8000/static/<filename>
	// route.PathPrefix("/web/static/css/").Handler(http.StripPrefix("/web/static/css/", http.FileServer(http.Dir(css))))
	// route.PathPrefix("/web/static/js/").Handler(http.StripPrefix("/web/static/js/", http.FileServer(http.Dir(js))))

	// route.PathPrefix("/web/static/css/").Handler(s).PathPrefix("/static/").Handler(s)
	// route.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	route.HandleFunc("/", controllers.GetAllTasks)
	route.HandleFunc("/add", controllers.Insert).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/complete/{id}", controllers.Complete)

	return route
}
