package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// Product model
type Product struct {
	//gorm.Model
	ID    uint   `gorm:"primary_key"`
	Code  string `gorm:"type:nvarchar(100);NOT NULL"`
	Price uint
}

func main() {
	db, err := gorm.Open("mssql", "server=127.0.0.1;port=1433;trusted_connection=yes;database=T1_DB")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//Drop Table
	//db.DropTableIfExists(&Product{})

	// Migrate the schema
	db.AutoMigrate(&Product{})

	//db.Table("Orders").CreateTable(&Product{})

	// a := &Product{
	// 	Code:  "L1214",
	// 	Price: 20000,
	// }

	// // Create
	// db.Create(a)                 //.Table("Orders")
	// db.Table("Orders").Create(a) //.Table("Orders")
	// db.Save(a)

	// // Update or Create
	// a := &Product{
	// 	ID:    1,
	// 	Code:  "L1214",
	// 	Price: 10800,
	// }
	// db.Save(a) //.Table("Orders")

	// // Read
	var product Product

	// db.First(&product, "price = ?", 20000) // find product with code l1212
	db.First(&product, 1) // find product with id 1

	// // // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	//db.Table("products").Where("code = ? and price = ? ", "L1214", 10800).Update("Code", "L1213")
	//db.Table("products").Where("code = ?", "L1213").Update("Code", "L1212")

	// // Delete - delete product
	db.Delete(&product)
}
