package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"enderz.net/testcontainer-test/internal/config"
	"enderz.net/testcontainer-test/internal/data"
	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/rest"
	_ "github.com/microsoft/go-mssqldb"
)

type application struct {
	config *config.Config
	db     *sql.DB
	models data.Models
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// prefer using a config file
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error config file:", err)
		os.Exit(1)
	}

	db, err := database.OpenDB(cfg.DB)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}

	dbTimeout := time.Duration(cfg.DB.Timeout) * time.Second
	app := &application{
		config: cfg,
		db:     db,
		models: data.NewModels(db, &dbTimeout),
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	fmt.Println("Starting server!")
	err = srv.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}

func (a *application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	rest.RespondWithJSON(w, r, http.StatusOK, "available", nil)
}
