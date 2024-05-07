package main

import (
	"flag"
	"log"
	"time"

	"syncer-monitor/prometh"
	"syncer-monitor/sqlexec"
)

var DSN = flag.String("dsn", "", "postgres dsn")
var pushAddr = flag.String("pushAddr", "", "pushAddr")
var interval = flag.Int("interval", 10, "check interval(min)")

func main() {
	flag.Parse()
	if *DSN == "" || *pushAddr == "" {
		log.Fatalln("dsn or pushAddr is empty")
	}

	db, err := sqlexec.InitDB(*DSN)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	for {
		syncers, err := sqlexec.Check(db)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(*interval) * time.Minute)
		}

		log.Println("syncers:", syncers)

		//TODO:push(syncers)
		for _, syncer := range syncers {
			err = prometh.Push(*pushAddr, syncer.Name, syncer.Epoch)
			if err != nil {
				log.Println(err)
			}

		}

		time.Sleep(time.Duration(*interval) * time.Minute)
	}

}
