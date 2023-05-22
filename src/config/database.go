package config

import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "syamsf/learning-gin/src/helper"
)

func DB() *gorm.DB {
  host 	 	 := "localhost"
  port 	 	 := "3306"
  dbName 	 := "belajar_golang_gin"
  username := "root"
  password := "password"

  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  helper.PanicIfError(err, "database")

  return db
}
