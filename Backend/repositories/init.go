package repositories

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user = "root"
	db_pwd  = "123456"
)

func InitMysql() {
	dsn := db_user + ":" + db_pwd + "@tcp(116.204.108.126:3306)/resume"
	db, err := sql.Open("mysql", dsn)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
