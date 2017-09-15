package main

import (
	"g"
	"log"
	"http"
	"flag"
)


func main() {
	//Prase the vars
	pconfName := flag.String("c", "C:/Data/WorkSpace/Go/awesomeProject/src/cfg.json", "configure file")
	flag.Parse()

	//Read config filezou
	err := g.ReadConf(*pconfName)
	if err != nil {
		log.Fatalf("ReadConf Err %s: %s", g.ConfFile(), err)
	}

	//start http
	http.Start()

}
