package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/infrastructure"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func ConfigureMappings(app mux.Router, handlers infrastructure.HandlerContainer) {
	//ping
	app.HandleFunc("/ping", handlers.Ping.Handle).Methods("GET")

	//material
	app.HandleFunc("/api/material", handlers.MaterialHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/material", handlers.MaterialHandler.Create).Methods("POST")
	app.HandleFunc("/api/material", handlers.MaterialHandler.Update).Methods("PUT")
	app.HandleFunc("/api/material/{id}", handlers.MaterialHandler.Delete).Methods("DELETE")

	//material type
	app.HandleFunc("/api/materialType", handlers.MaterialTypeHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/materialType", handlers.MaterialTypeHandler.Create).Methods("POST")
	app.HandleFunc("/api/materialType/units", handlers.MaterialTypeHandler.GetUnitsOfMeasurement).Methods("GET")
	app.HandleFunc("/api/materialType", handlers.MaterialTypeHandler.Update).Methods("PUT")
	app.HandleFunc("/api/materialType/{id}", handlers.MaterialTypeHandler.Delete).Methods("DELETE")

	//movement
	app.HandleFunc("/api/movement", handlers.MovementHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/movement/{type}", handlers.MovementHandler.GetAllByType).Methods("GET")
	app.HandleFunc("/api/movement/byId/{id}", handlers.MovementHandler.GetById).Methods("GET")
	/*app.HandleFunc("/api/movement", handlers.MovementHandler.Create).Methods("POST")
	app.HandleFunc("/api/movement", handlers.MovementHandler.Update).Methods("PUT")*/

	//user
	app.HandleFunc("/api/user/signup", handlers.UserHandler.Signup).Methods("POST")
	app.HandleFunc("/api/user/login", handlers.UserHandler.Login).Methods("POST")
	app.HandleFunc("/api/user/logout", handlers.UserHandler.Logout).Methods("POST")
	app.HandleFunc("/api/user/{id}", handlers.UserHandler.Logout).Methods("PUT")
	app.HandleFunc("/api/user/activeDesactiveUser/{id}", handlers.UserHandler.ActiveDesactiveUser).Methods("PUT")
	app.HandleFunc("/api/user", handlers.UserHandler.Update).Methods("PUT")

	//employee
	app.HandleFunc("/api/employee", handlers.EmployeeHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/employee/{id}", handlers.EmployeeHandler.GetById).Methods("GET")
	app.HandleFunc("/api/employee", handlers.EmployeeHandler.Update).Methods("PUT")
	//charge
	app.HandleFunc("/api/charge", handlers.ChargeHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/charge", handlers.ChargeHandler.Create).Methods("POST")

	//product
	app.HandleFunc("/api/product", handlers.ProductHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/product/byName", handlers.ProductHandler.GroupedByName).Methods("GET")
	app.HandleFunc("/api/product/byNameMap", handlers.ProductHandler.GroupedByNameMap).Methods("GET")
	app.HandleFunc("/api/product/byNameMapWithStock", handlers.ProductHandler.GroupedByNameMapWithStock).Methods("GET")
	app.HandleFunc("/api/product/{id}", handlers.ProductHandler.GetById).Methods("GET")
	app.HandleFunc("/api/product/name/{name}", handlers.ProductHandler.GetByName).Methods("GET")
	app.HandleFunc("/api/product", handlers.ProductHandler.Create).Methods("POST")
	app.HandleFunc("/api/product", handlers.ProductHandler.Update).Methods("PUT")
	app.HandleFunc("/api/product/{id}", handlers.ProductHandler.Delete).Methods("DELETE")
	app.HandleFunc("/api/product/assignation", handlers.ProductHandler.AssignMaterialsToProduct).Methods("POST")

	//movement
	app.HandleFunc("/api/movement", handlers.MovementHandler.Create).Methods("POST")

	fmt.Println("Starting app in localhost:8080")
	handler := cors.AllowAll().Handler(&app)
	log.Fatal(http.ListenAndServe(":8080", handler))

}
