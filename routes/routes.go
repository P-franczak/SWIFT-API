package routes

import (
    "github.com/gorilla/mux"
    "github.com/P-franczak/swift-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/v1/swift-codes/{swiftCode}", handlers.GetSwiftCode).Methods("GET")
    router.HandleFunc("/v1/swift-codes/country/{countryISO2}", handlers.GetSwiftCodesByCountry).Methods("GET")
    router.HandleFunc("/v1/swift-codes", handlers.CreateSwiftCode).Methods("POST")
    router.HandleFunc("/v1/swift-codes/{swiftCode}", handlers.DeleteSwiftCode).Methods("DELETE")
}
