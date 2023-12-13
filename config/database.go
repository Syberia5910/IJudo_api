package config

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "judoApi/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("tounoi.db"), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to database!")
    }

    err = database.AutoMigrate(&models.Club{}, &models.Judoka{},
      &models.Tournament{}, &models.Category{}, &models.Tatami{},
      &models.Pool{}, &models.Registration{})
    if err != nil {
        return
    }

    DB = database
}