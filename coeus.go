package coeus

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//Database Default Config Struct is bottom of the file
var DB *sql.DB

func Construct(dbConfig Config) bool {
	//Getting Database Config

	//Setting connection string
	var connectionString string
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	//Starting Database Connection
	var err error
	DB, err = sql.Open("mysql", connectionString)

	DB.SetConnMaxLifetime(time.Minute * 1)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(5)

	//Database Connection Error Handling
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

type Config struct {
	Type     string //mysql mssql
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Instance string //For app engine CloudSQL
}