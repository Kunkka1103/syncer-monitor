package sqlexec

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Syncer struct {
	Name  string
	Epoch int64
}

func InitDB(DSN string) (DB *sql.DB, err error) {

	DB, err = sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}

	info := fmt.Sprintf("dsn check success")
	log.Println(info)

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	info = fmt.Sprintf("database connect success")
	log.Println(info)

	return DB, nil
}

func Check(db *sql.DB) (syncers []Syncer, err error) {
	SQL := "SELECT * FROM chain.sync_syncers WHERE name NOT LIKE '%test%';"
	rows, err := db.Query(SQL)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var s Syncer

	for rows.Next() {
		err = rows.Scan(&s.Name, &s.Epoch)
		if err != nil {
			return nil, err
		}
		syncers = append(syncers, s)
	}
	return syncers, nil
}
