package main

import (
	"flag"
	"log"
	"time"

	"syncer-monitor/prometh"
	"syncer-monitor/sqlexec"
)

var DSN = flag.String("dsn", "", "postgres dsn")
var pushAddr = flag.String("push-addr", "", "Address of the Pushgateway to send metrics")
var interval = flag.Int("interval", 10, "Interval in minutes to check the URLs")

func main() {
	flag.Parse()
	if *DSN == "" || *pushAddr == "" {
		log.Fatalln("dsn or push-addr is empty")
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

		height := CalculateHeight()
		//TODO:push(syncers)
		for _, syncer := range syncers {
			delay := height - syncer.Epoch
			err = prometh.Push(*pushAddr, syncer.Name, delay)
			if err != nil {
				log.Println(err)
			}

		}

		time.Sleep(time.Duration(*interval) * time.Minute)
	}

}

func CalculateHeight() int64 {
	// 设置区块链初始时间（高度0的时间）
	location, _ := time.LoadLocation("Asia/Shanghai")
	startTime := time.Date(2020, 8, 25, 6, 0, 0, 0, location)

	// 获取当前时间
	currentTime := time.Now().In(location)

	// 计算时间差（以秒为单位）
	timeDifference := currentTime.Sub(startTime).Seconds()

	// 计算当前高度
	currentHeight := int64(timeDifference / 30)

	return currentHeight
}
