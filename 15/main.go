package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"toplearn/15/router"
)

type config struct {
	Server string `json:"server"`
	Lang   string `json:"lang"`
}

func main() {
	fmt.Println("open file for config file")
	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Printf("error : %s", err)
		log.Fatalln("error : ", err)
	}
	fmt.Println("decode json file")
	conf := new(config)
	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		fmt.Printf("error : %s", err)
		log.Fatalln("error : ", err)
	}
	fmt.Println("starting server for config file ")
	router.Rooter(conf.Server,conf.Lang)

}
