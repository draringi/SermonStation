package db

import (
	"encoding/json"
	"errors"
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

func (r *Recording) save() error {
	if r.rID > 0 {
		resp, err := connection.query("UPDATE recordings SET recorded_at = ?, title = ?, preacher = ?, path = ? WHERE rid = ?", r.at, r.title, r.preacher.pID, r.path, r.rID)
		if err != nil {
			return err
		}
		resp.Close()
	} else {
		resp, err := connection.query("INSERT INTO recordings (recorded_at, title, preacher, path) VALUES (?, ?, ?, ?) RETURNING pid", r.at, r.title, r.preacher.pID, r.path)
		if err != nil {
			return err
		}
		defer resp.Close()
		if !resp.Next() {
			return errors.New("Something went wrong parsing the responce. Expected 1 row, got none.")
		}
		err = resp.Scan(r.rID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Recording) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RID      int        `json:"id"`
		Title    string     `json:"title"`
		At       *time.Time `json:"at"`
		Preacher *Preacher  `json:"preacher"`
		Path     string     `json:"path"`
	}{
		r.rID,
		r.title,
		r.at,
		r.preacher,
		r.path,
	})
}

func NewRecording(at *time.Time, path string, preacher *Preacher) (*Recording, error) {
	r := new(Recording)
	r.at = at
	r.path = path
	r.preacher = preacher
	err := r.save()
	if err != nil {
		return nil, err
	}
	return r, nil
}
