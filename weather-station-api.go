package main

import (
	"os"
	"log"
    "bytes"
	"net/http"
	"database/sql"

    "github.com/fiskeben/weather-station-api/context"

	_ "github.com/lib/pq"
)


func main() {
    if len(os.Args) < 2 {
        panic("Please specify path to configuration file")
    }

	var err error

    configuration := LoadConfiguration(os.Args[1])
    connectionString := getConnectionString(configuration)
    log.Printf(connectionString)

	db, err := sql.Open("postgres", connectionString)
	
	if err != nil {
		log.Fatalf("Unable to initialize database connection: %s", err.Error())
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err.Error())
		os.Exit(2)
	}

    appContext := context.AppContext{Db: db}

	router := NewRouter(&appContext)
    log.Fatal(http.ListenAndServe(":" + configuration.Http.Port, router))
}

func getConnectionString(config Configuration) string {
    var buffer bytes.Buffer

    buffer.WriteString("user=")
    buffer.WriteString(config.Database.Username)
    buffer.WriteString(" password=")
    buffer.WriteString(config.Database.Password)
    buffer.WriteString(" dbname=")
    buffer.WriteString(config.Database.DatabaseName)
    buffer.WriteString(" sslmode=disable")

    return buffer.String()
}