package models

import "gorm.io/gorm"

type SwiftCode struct {
    gorm.Model
    SwiftCode     string `gorm:"uniqueIndex;not null"`
    BankName      string `gorm:"not null"`
    Address       string
    CountryISO2   string `gorm:"not null"`
    CountryName   string `gorm:"not null"`
    IsHeadquarter bool
}
