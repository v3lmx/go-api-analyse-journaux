package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/v3lmx/go-api-analyse-journaux/internal/core"
)

const apiPort = "15442"

var ErrUnsupportedMediaType = fmt.Errorf("Unsupported media type")
var ErrMissingRequiredHeader = fmt.Errorf("Missing required header: 'Accept'")

func Start(logger *log.Logger, journalService core.JournalService) {
	http.Handle("/analysis", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		analysisHandler(w, r, journalService, logger)
	}))

	logger.Println("Starting server on port", apiPort)
	logger.Fatal(http.ListenAndServe(":"+apiPort, nil))
}

func analysisHandler(w http.ResponseWriter, r *http.Request, journalService core.JournalService, logger *log.Logger) {
	if err := validateAnalysisRequest(r); err != nil {
		switch err {
		case ErrUnsupportedMediaType:
			http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		case ErrMissingRequiredHeader:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		logger.Println("ERROR: Failed to process analysis request:", err)
		return
	}

	report, err := journalService.Analyse()
	if err != nil {
		logger.Println("ERROR: Failed to process analysis request:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toResponse(w, report)
	logger.Println("Analysis request successfully processed")
}

func validateAnalysisRequest(r *http.Request) error {
	acceptHeader := r.Header.Get("Accept")
	if acceptHeader == "" {
		return ErrMissingRequiredHeader
	}
	if acceptHeader != "text/csv" {
		return ErrUnsupportedMediaType
	}

	return nil
}

func toResponse(w http.ResponseWriter, report core.JournalReport) {
	w.Header().Set("Content-Type", "text/csv")
	for _, entry := range report {
		_, err := w.Write([]byte(fmt.Sprintf("%s,%s,%s,%s\n", entry.Day, entry.Hour, entry.File, entry.Message)))
		if err != nil {
			err = fmt.Errorf("failed to write response: %w", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
