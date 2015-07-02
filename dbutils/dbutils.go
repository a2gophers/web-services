package dbutils

import (
	"bytes"
	"database/sql"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// input flags
	dbHost  = flag.String("ws-db-host", "", "database host to connect to")
	dbPort  = flag.String("ws-db-port", "", "database port to connect on")
	dbUser  = flag.String("ws-db-user", "", "user to connect to database as")
	dbPsswd = flag.String("ws-db-psswd", "", "user password to authenticate with")
	dbDB    = flag.String("ws-db-db", "", "database to connect to on host")

	// option flags
	dontParseTime = flag.Bool("ws-db-dont-parse-time", true, "don't parse time to time.Time")
)

type DBInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string

	ParseTime bool
}

// dbURI builds the appropriate dbURI from the given db info
// NOTE: it is dumb, it will not be intelligent, that is up to the
// user to check.
func dbURI(info DBInfo) string {
	var uri = bytes.NewBufferString(info.User)
	uri.WriteString(":")
	uri.WriteString(info.Password)
	uri.WriteString("@")
	uri.WriteString(info.Host)
	if len(info.Port) > 0 {
		uri.WriteString(":")
		uri.WriteString(info.Password)
	}

	uri.WriteString("/")
	uri.WriteString(info.DB)
	if info.ParseTime {
		uri.WriteString("?parseTime=true")
	}

	return uri.String()
}

func DBURIFromFlags() DBInfo {
	return DBInfo{
		User:      *dbUser,
		Password:  *dbPsswd,
		Host:      *dbHost,
		Port:      *dbPort,
		DB:        *dbDB,
		ParseTime: *dontParseTime,
	}
}

// DBConnFromFlags initializes a MySQL connection from
// command line flags.
func DBConnFromFlags() (*sql.DB, error) {
	return sql.Open("mysql", dbURI(DBURIFromFlags()))
}
