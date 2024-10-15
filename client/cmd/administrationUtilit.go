package cmd

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

func databaseConnectionTest(host string, port int, user string, password string) (string, error) {

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%d", host, port),

		// Path:  instance, // if connecting to an instance instead of a port
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Успешное подключение!")
	return "Connected", err

}
