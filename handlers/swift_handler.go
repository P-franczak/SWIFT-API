package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/P-franczak/swift-api/database"

	"github.com/P-franczak/swift-api/models"
	"github.com/gorilla/mux"
)

// Pobieranie pojedynczego kodu SWIFT
func GetSwiftCode(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var swiftCode models.SwiftCode
    database.DB.First(&swiftCode, "swift_code = ?", params["swiftCode"])

    if swiftCode.SwiftCode == "" {
        http.Error(w, "SWIFT code not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(swiftCode)
}

// Pobieranie wszystkich kodów SWIFT dla danego kraju
func GetSwiftCodesByCountry(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var swiftCodes []models.SwiftCode
    database.DB.Where("country_iso2 = ?", params["countryISO2"]).Find(&swiftCodes)
    json.NewEncoder(w).Encode(swiftCodes)
}

// Dodawanie nowego SWIFT kodu
func CreateSwiftCode(w http.ResponseWriter, r *http.Request) {
    var swiftCode models.SwiftCode
    json.NewDecoder(r.Body).Decode(&swiftCode)
    database.DB.Create(&swiftCode)
    json.NewEncoder(w).Encode(map[string]string{"message": "SWIFT code added successfully"})
}

// Usuwanie SWIFT kodu
func DeleteSwiftCode(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    database.DB.Delete(&models.SwiftCode{}, "swift_code = ?", params["swiftCode"])
    json.NewEncoder(w).Encode(map[string]string{"message": "SWIFT code deleted successfully"})
}
