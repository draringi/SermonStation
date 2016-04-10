package db

import (
	"time"
)

type Recording struct {
	rID      int
	at       *time.Time
	title    string
	preacher *Preacher
	path     string
}

func (r *Recording) ID() int {
	return r.rID
}

func (r *Recording) At() *time.Time {
	return r.at
}

func (r *Recording) Title() string {
	return r.title
}

func (r *Recording) Preacher() *Preacher {
	return r.preacher
}

func (r *Recording) Path() string {
	return r.path
}

func newRecording(at *time.Time, path string) *Recording {
	return nil
}
