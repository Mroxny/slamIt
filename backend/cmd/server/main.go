package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/Mroxny/slamIt/docs"
	"github.com/Mroxny/slamIt/internal/router"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			SlamIt API
//	@version		1.0
//	@description	API for managing poetry slams and participation.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	support@slamit.app

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1
//	@schemes	http

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	testData := flag.Bool("test-data", false, "Start the server instance with some test data")
	flag.Parse()

	r := chi.NewRouter()
	var routeHandler http.Handler

	if *testData {
		routeHandler = router.SetupTestRouter()
	} else {
		routeHandler = router.SetupV1Router()
	}

	r.Mount("/api/v1", routeHandler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
