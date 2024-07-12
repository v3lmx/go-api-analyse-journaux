package repository

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/v3lmx/go-api-analyse-journaux/internal/core"
)

type FileJournalRepository struct {
	FilePath string
}

func NewFileJournalRepository(filePath string) FileJournalRepository {
	return FileJournalRepository{filePath}
}

func (r FileJournalRepository) GetJournal() (core.Journal, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		err = fmt.Errorf("failed to open file: %w", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	records := make([][]string, 0)
	for scanner.Scan() {
		// Some messages have commas in them,
		// so we put everything after the second comma in the message field.
		values := strings.SplitN(scanner.Text(), ",", 3)
		records = append(records, values)
	}

	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("failed to read file: %w", err)
		return core.Journal{}, err
	}

	var journal core.Journal
	for _, r := range records {
		date, err := time.Parse(time.RFC3339, r[0])
		if err != nil {
			err = fmt.Errorf("failed to parse date: %w", err)
			return nil, err
		}
		journal = append(journal, core.JournalEntry{
			Date:    date,
			File:    r[1],
			Message: r[2],
		})
	}

	return journal, nil

}
