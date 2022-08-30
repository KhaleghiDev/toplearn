package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type register struct {
	ID         int
	username   string
	email      string
	password   string
	creates_at string
	ip         string
}

func main() {
	fmt.Println("mysql db :")
	db, err := sql.Open("mysql", "root:root@/eventdb")
	rows, err := db.Query("SELECT * FROM eventdb.register")
	if err != nil {
		log.Fatal(err)
	}
	reg := []register{}
	for rows.Next() {
		r := register{}
		err = rows.Scan(&r.ID, &r.username, &r.email, &r.password, &r.creates_at, &r.ip)
		if err != nil {
			log.Println(err)
			continue
		}
		reg = append(reg, r)
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(reg)
	}
	fmt.Println(db.Stats())
	//db.Exec("INSERT INTO `eventdb`.`register` (`username`, `email`, `password`, `createat_at`, `definelogin`) VALUES (?, ?, '12345678', '2019-07-01', '192.168.2.2'),"amir","amir@gmail.com")
	
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
