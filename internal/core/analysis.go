package core

import (
	"fmt"
)

func (svc JournalService) Analyse() (JournalReport, error) {
	journal, err := svc.journalRepository.GetJournal()
	if err != nil {
		err = fmt.Errorf("failed to get journal: %w", err)
		return JournalReport{}, err
	}

	report, err := journal.Report()
	if err != nil {
		err = fmt.Errorf("failed to generate report: %w", err)
		return JournalReport{}, err
	}

	return report, nil
}

type JournalService struct {
	journalRepository JournalRepository
}

func NewJournalService(journalRepository JournalRepository) JournalService {
	return JournalService{journalRepository}
}

type JournalRepository interface {
	GetJournal() (Journal, error)
}
