// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

package main

import (
	"avvio-api/docs"
	_ "avvio-api/docs"
	route_handlers "avvio-api/route-handlers"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	var api = r.PathPrefix("/api").Subrouter()
	var api1 = api.PathPrefix("/v1").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	/*
	//Use to enable authentication
	api.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("x-auth-token") != "admin" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})
	*/
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	docs.SwaggerInfo.Title = "AVVIO Swagger API"
	docs.SwaggerInfo.Description = "This is the avvio API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	httpSwagger.URL("/docs/doc.json")
	r.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	api1.HandleFunc("/application/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/application", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/application", route_handlers.CreateTaskHandler).Methods( "POST")

	api1.HandleFunc("/project/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/project", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/project", route_handlers.CreateTaskHandler).Methods( "POST")

	api1.HandleFunc("/task/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/task", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/task", route_handlers.CreateTaskHandler).Methods( "POST")

	api1.HandleFunc("/issue/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/issue", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/issue", route_handlers.CreateTaskHandler).Methods( "POST")

	api1.HandleFunc("/pipeline/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/pipeline", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/pipeline", route_handlers.CreateTaskHandler).Methods( "POST")

	api1.HandleFunc("/team/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/team", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/team", route_handlers.CreateTaskHandler).Methods( "POST")

	api1.HandleFunc("/team_member/{id}", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/team_member", route_handlers.GetTaskHandler).Methods("GET")
	api1.HandleFunc("/team_member", route_handlers.CreateTaskHandler).Methods( "POST")
	return r
}

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
