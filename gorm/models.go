package main

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    Name      string
    Age       uint8
    Birthdate time.Time
}
