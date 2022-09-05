package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
func connMysql()(*gorm.DB ,error){
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return  db,err
}
func connsqlite()(*gorm.DB ,error){
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	return  db,err
}
func main() {
db,err:=connMysql()
// db,err:=connsqlite()
if err != nil {
	panic("failed to connect database")
}
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "D43", Price: 200})
	db.Create(&Product{Code: "D44", Price: 150})
	// Read
	var product Product
	db.First(&product, 5)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&Product{}, []int{7,8,9})
}
