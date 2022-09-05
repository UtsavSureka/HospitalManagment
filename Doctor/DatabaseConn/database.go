package DatabaseConn

import (
	"database/sql"
	"log"
)

var Dbconn *sql.DB

func SetUpDBConn() {
	pass := "root:Ram@161198@/doctor"
	var err error
	Dbconn, err = sql.Open("mysql", pass)
	if err != nil {
		log.Fatal(err)
	}
}
