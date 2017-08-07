package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Link struct {
  ID        uint
  Url       string
  Title     string
  Category  string
  Done      bool
}

func main() {
  db, err := gorm.Open("sqlite3", "links.db")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Link{})

  // Create
  db.Create(&Link{Url: "http://someurl.to.site", Title: "Javascript Page", Category: "Javascript", Done: 1})

  // Read
  var link Link
  db.First(&link, 1) // find link with id 1
  db.First(&link, "category = ?", "Javascript") // find link with category Javascript

  // Update - update link's done to 0
  db.Model(&link).Update("Done", 0)

  // Delete - delete link
  db.Delete(&link)
}
