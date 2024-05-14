package dbstuff

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// docker pull mysql
// docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql

func ConnectDB() {
	fmt.Println("Starting database")
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
