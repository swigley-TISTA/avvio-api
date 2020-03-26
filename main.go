// Testing go-swagger generation
//
// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Daniel<danielfs.ti@gmail.com>
//
//     Consumes:
//     - text/plain
//
//     Produces:
//     - text/plain
//
// swagger:meta
package main

import (
	"avvio-api/models"
	"avvio-api/route-handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tidwall/buntdb"
	"log"
	"net/http"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	r.HandleFunc("/swagger.json", swagger).Methods("GET")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	//r.HandleFunc("/team", route_handlers.GetTeamsHandler).Methods("GET")
	r.HandleFunc("/team", route_handlers.CreateTeamHandler).Methods("POST")
	r.HandleFunc("/team/{id}", route_handlers.GetTeamHandler).Methods("GET")

	r.HandleFunc("/task/{id}", route_handlers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", route_handlers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", route_handlers.CreateTaskHandler).Methods("POST")
	//r.HandleFunc("/application/{id}", )
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	models.InitStore(db)
	http.ListenAndServe(":8080", r)
}

// swagger:operation GET /hello/{name} hello Hello
//
// Returns a simple Hello message
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - name: name
//   in: path
//   description: Name to be returned.
//   required: true
//   type: string
// responses:
//   '200':
//     description: The hello message
//     type: string
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "swagger.json")
}

