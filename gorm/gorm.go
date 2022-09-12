package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

func main() {
    dsn := "root:admin@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }
    db.AutoMigrate(&User{})
    // Create Record
    user := User{Name: "Qing Wang", Age: 18, Birthdate: time.Now()}
    result := db.Create(&user)
    fmt.Println("%d records created", result.RowsAffected)
}
