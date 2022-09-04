package main

import (
	"database/sql"
	"fmt"

	"log"

	// Import the Azure AD driver module (also imports the regular driver package)
	_ "github.com/denisenkom/go-mssqldb"
)

type register struct {
	ID    int
	Name  string
	Fname string
}

func main() {
	fmt.Println("mysql db :")
	// connectDB()
	// metod all
	all()
	//metod show
	show(2)
	// // metod insert
	insert("hosin","hoseni")
	all()
	// //updete
	update(1, "nima")
	show(1)
	// delete(1)
	delete(3)
}
func connectDB() (*sql.DB, error) {
	nameDriver := "mssql"
	conneact := "server=127.0.0.1;port=1433;trusted_connection=yes;database=dbs_test"
	db, err := sql.Open(nameDriver, conneact)
	return db, err
}
func update(id int, name string) error {
	db, err := connectDB()
	reu, err := db.Exec("UPDATE tbl_testUser SET name = ? WHERE (id = ?)", name, id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reu.LastInsertId())
	fmt.Println(reu.RowsAffected())
	return err

}
func delete(id int) error {
	db, err := connectDB()
	red, err := db.Exec("DELETE FROM tbl_testUser ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(red.LastInsertId())
	fmt.Println(red.RowsAffected())
	return err
}
func insert(name string, fname string) (error,sql.Result)  {
	db, err := connectDB()
	res, err := db.Exec("INSERT INTO tbl_testUser (name,fname) VALUES (?, ?)", name,fname)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(res.LastInsertId())
	// fmt.Println(res.RowsAffected())

	return err ,res
}
func show(id int) error {
	db, err := connectDB()
	row := db.QueryRow("SELECT * FROM tbl_testUser where id=?", id)
	r := register{}
	err = row.Scan(&r.ID, &r.Name,&r.Fname)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(r)
	return err
}
func all() error {
	db, err := connectDB()
	rows, err := db.Query("SELECT * FROM tbl_testUser")
	if err != nil {
		log.Fatal(err)
	}
	reg := []register{}
	for rows.Next() {
		r := register{}
		err = rows.Scan(&r.ID, &r.Name,&r.Fname)
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
