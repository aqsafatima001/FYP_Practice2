package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

// config is the type for all application configuration
type config struct {
	port int // what port do we want the web server to listen on
}

// application is the type for all data we want to share with the
// various parts of our application. We will share this information in most
// cases by using this type as the receiver for functions
type application struct {
	config   config
	db       *sql.DB
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewApplication(config config, db *sql.DB, infoLog, errorLog *log.Logger) *application {
	return &application{
		config:   config,
		db:       db,
		infoLog:  infoLog,
		errorLog: errorLog,
	}
}

func main() {
	// Connection with Database
	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=learning;user id=Final_Year_Project;password=fyp;"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err.Error())
	}
	fmt.Println("Connected to the database!")

	// Initialize application
	var cfg config
	cfg.port = 8090

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := NewApplication(cfg, db, infoLog, errorLog)

	// Serve
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
