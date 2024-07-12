package core

import (
	"fmt"
	"time"
)

var ErrorNoUniqueMax = fmt.Errorf("Error: there is no unique maximum entry.")
var ErrorEmptyJournal = fmt.Errorf("Error: journal is empty.")

type JournalReport []ReportEntry

type ReportEntry struct {
	Day     string
	Hour    string
	File    string
	Message string
}

type ReportTime struct {
	Day  string
	Hour string
}

type ReportValues struct {
	File    string
	Message string
}

type Reporter interface {
	Report() JournalReport
}

func (journal Journal) Report() (JournalReport, error) {
	var report JournalReport

	if len(journal) == 0 {
		return JournalReport{}, ErrorEmptyJournal
	}

	var maxEntry ReportEntry
	var maxCount int
	var maxExists bool
	count := make(map[ReportValues]int)

	currentTime := ReportTime{dateToDay(journal[0].Date), dateToHour(journal[0].Date)}
	// IMPROVEMENT: refactor this loop to be more readable and maintainable.
	for _, entry := range journal {
		entryTime := ReportTime{dateToDay(entry.Date), dateToHour(entry.Date)}
		// if these values are different, we move into the next
		// time segment, so we have to add the max entry to the
		// report and reset the count.
		if entryTime != currentTime {
			if !maxExists {
				return JournalReport{}, ErrorNoUniqueMax
			}
			report = append(report, maxEntry)
			currentTime = entryTime
			count = make(map[ReportValues]int)
			maxCount = 0
		}

		// update loop variables with the current entry
		reportValues := ReportValues{entry.File, entry.Message}
		count[reportValues]++
		currentCount := count[reportValues]

		// if we reach the maxCount, there are multiple entries
		// with the same count, so we can't have a unique max
		if currentCount == maxCount {
			maxExists = false
		}

		// if we go over maxCount, we have a new max,
		// so we can reset the maxExists flag
		if currentCount > maxCount {
			maxExists = true
			maxCount = currentCount
			// IMPROVEMENT: this does not need to be
			// rebuilt if it is already the same
			// values
			maxEntry = ReportEntry{
				Day:     dateToDay(entry.Date),
				Hour:    dateToHour(entry.Date),
				File:    entry.File,
				Message: entry.Message,
			}
		}
	}
	// Handle the last entry
	// (this should not be necessary if the loop is refactored)
	if !maxExists {
		return JournalReport{}, ErrorNoUniqueMax
	}
	report = append(report, maxEntry)

	return report, nil
}

func dateToDay(date time.Time) string {
	return date.Format("01022006")
}

func dateToHour(date time.Time) string {
	return date.Format("15")
}
