package core

import "time"

type Journal []JournalEntry

type JournalEntry struct {
	Date    time.Time
	File    string
	Message string
}
