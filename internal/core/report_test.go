package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newJournalEntry(date time.Time, file, message string) JournalEntry {
	return JournalEntry{
		Date:    date,
		File:    file,
		Message: message,
	}
}

func TestReportEmptyJournal(t *testing.T) {
	t.Parallel()

	j := Journal{}
	r, err := j.Report()

	assert.Empty(t, r)
	assert.EqualError(t, err, ErrorEmptyJournal.Error())
}

func TestReportOneLine(t *testing.T) {
	t.Parallel()

	j := Journal{{
		Date:    time.Date(2018, time.January, 02, 12, 01, 02, 00, time.UTC),
		File:    "file.go",
		Message: "Error: It worked on my machine",
	}}

	expectedReport := JournalReport{{
		Day:     "01022018",
		Hour:    "12",
		File:    "file.go",
		Message: "Error: It worked on my machine",
	}}

	r, err := j.Report()

	assert.Equal(t, expectedReport, r)
	assert.NoError(t, err)
}

func TestReportMultipleLines(t *testing.T) {
	t.Parallel()

	j := Journal{
		{
			Date:    time.Date(2018, time.January, 02, 12, 01, 02, 00, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 12, 01, 02, 01, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 12, 01, 02, 02, time.UTC),
			File:    "other.go",
			Message: "This is fine",
		},
	}

	expectedReport := JournalReport{{
		Day:     "01022018",
		Hour:    "12",
		File:    "file.go",
		Message: "Error: It worked on my machine",
	}}

	r, err := j.Report()

	assert.Equal(t, expectedReport, r)
	assert.NoError(t, err)
}

func TestReportMultiplesHours(t *testing.T) {
	t.Parallel()

	j := Journal{
		{
			Date:    time.Date(2018, time.January, 02, 12, 01, 02, 00, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 12, 01, 02, 01, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 12, 01, 02, 02, time.UTC),
			File:    "other.go",
			Message: "This is fine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 13, 01, 02, 02, time.UTC),
			File:    "other.go",
			Message: "This is fine",
		},
	}

	expectedReport := JournalReport{{
		Day:     "01022018",
		Hour:    "12",
		File:    "file.go",
		Message: "Error: It worked on my machine",
	}, {
		Day:     "01022018",
		Hour:    "13",
		File:    "other.go",
		Message: "This is fine",
	}}

	r, err := j.Report()

	assert.Equal(t, expectedReport, r)
	assert.NoError(t, err)
}

func TestReportHourNextDay(t *testing.T) {
	t.Parallel()

	j := Journal{
		{
			Date:    time.Date(2018, time.January, 02, 23, 01, 02, 00, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 23, 01, 02, 01, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 23, 01, 02, 02, time.UTC),
			File:    "other.go",
			Message: "This is fine",
		},
		{
			Date:    time.Date(2018, time.January, 03, 00, 01, 02, 02, time.UTC),
			File:    "other.go",
			Message: "This is fine",
		},
	}

	expectedReport := JournalReport{{
		Day:     "01022018",
		Hour:    "23",
		File:    "file.go",
		Message: "Error: It worked on my machine",
	}, {
		Day:     "01032018",
		Hour:    "00",
		File:    "other.go",
		Message: "This is fine",
	}}

	r, err := j.Report()

	assert.Equal(t, expectedReport, r)
	assert.NoError(t, err)
}

func TestReportNoUnique(t *testing.T) {
	t.Parallel()

	j := Journal{
		{
			Date:    time.Date(2018, time.January, 02, 10, 01, 02, 00, time.UTC),
			File:    "file.go",
			Message: "Error: It worked on my machine",
		},
		{
			Date:    time.Date(2018, time.January, 02, 10, 01, 02, 01, time.UTC),
			File:    "other.go",
			Message: "This is fine",
		},
	}
	r, err := j.Report()

	assert.Empty(t, r)
	assert.EqualError(t, err, ErrorNoUniqueMax.Error())
}
