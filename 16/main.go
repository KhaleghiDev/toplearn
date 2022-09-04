package main

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type register struct {
	ID         int
	Name       string
	email      string
	password   string
	creates_at string
	ip         string
}

func main() {
	fmt.Println("mysql db :")
	// metod all
	all()
	//metod show
	show(5)
	// metod insert
	insert("hosine")
	//updete
	update(4, "nima")
	delete(1)
}
func connectDB() (*sql.DB, error) {
	nameDriver := "mysql"
	conneact := "root:@/peple"
	db, err := sql.Open(nameDriver, conneact)
	return db, err
}
func update(id int, name string) error {
	db, err := connectDB()
	reu, err := db.Exec("UPDATE `peple`.`persian` SET `persianname` = ? WHERE (`Id` = ?);", name, id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reu.LastInsertId())
	fmt.Println(reu.RowsAffected())
	return err

}
func delete(id int) error {
	db, err := connectDB()
	red, err := db.Exec("DELETE FROM `persian` WHERE (`Id` = ?)", id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(red.LastInsertId())
	fmt.Println(red.RowsAffected())
	return err
}
func insert(name string) error {
	db, err := connectDB()
	res, err := db.Exec("INSERT INTO `persian` (`persianname`) VALUES (?)", name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
	return err
}
func show(id int) error {
	db, err := connectDB()
	row := db.QueryRow("SELECT * FROM peple.persian where id=?", id)
	r := register{}
	err = row.Scan(&r.ID, &r.Name)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(r)
	return err
}
func all() error {
	db, err := connectDB()
	rows, err := db.Query("SELECT * FROM peple.persian;")
	if err != nil {
		log.Fatal(err)
	}
	reg := []register{}
	for rows.Next() {
		r := register{}
		err = rows.Scan(&r.ID, &r.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		reg = append(reg, r)
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}

	}
	fmt.Println(reg)
	return err
}
