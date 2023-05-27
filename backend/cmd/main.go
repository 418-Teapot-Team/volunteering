package main

import (
	"github.com/BoryslavGlov/logrusx"
	"github.com/subosito/gotenv"
	"log"
	"os"
	app "volunteering/internal/app"
	"volunteering/pkg/repository"
)

func main() {

	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	logg, err := logrusx.New("budget-tracker")
	if err != nil {
		log.Fatal("error while trying to create logg instance")
	}

	db, err := repository.NewDB(os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatal("error while trying to connect to database", err)
	}

	repo := repository.NewSQL(db)

	api := app.NewApp(logg, repo)

	app.StartServer(api, os.Getenv("PORT"))

}
