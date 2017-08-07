package main

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/dialects/mysql"
)

type Link struct {
  ID        uint
  Url       string
  Title     string
  Category  string
  Done      bool
}

fun main() {
  db,err := gorm.Open("mysql", "jason:@/saved?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  db.CreateTable(&Link{})

  println("done")
}
