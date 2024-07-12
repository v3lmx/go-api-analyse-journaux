package main

import (
	"log"
	"os"

	"github.com/v3lmx/go-api-analyse-journaux/internal/api"
	"github.com/v3lmx/go-api-analyse-journaux/internal/core"
	"github.com/v3lmx/go-api-analyse-journaux/internal/repository"
)

func main() {
	logger := log.New(os.Stdout, "analysis: ", log.Flags())

	journalRepository := repository.NewFileJournalRepository("testdata/journal.csv")
	journalService := core.NewJournalService(journalRepository)

	api.Start(logger, journalService)
}
