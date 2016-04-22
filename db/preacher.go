package db

import (
	"encoding/json"
	"errors"
)

type Preacher struct {
	pID  int
	name string
}

func (p *Preacher) ID() int {
	return p.pID
}

func (p *Preacher) Name() string {
	return p.name
}

func newPreacher(name string) *Preacher {
	p := new(Preacher)
	p.name = name
	p.save()
	return p
}

func (p *Preacher) save() error {
	if p.pID > 0 {
		resp, err := connection.query("UPDATE preachers SET name = ? WHERE pid = ?", p.name, p.pID)
		if err != nil {
			return err
		}
		resp.Close()
	} else {
		resp, err := connection.query("INSERT INTO preachers (name) VALUES (?) RETURNING pid", p.name)
		if err != nil {
			return err
		}
		defer resp.Close()
		if !resp.Next() {
			return errors.New("Something went wrong parsing the responce. Expected 1 row, got none.")
		}
		err = resp.Scan(p.pID)
		if err != nil {
			return err
		}
	}
	return nil
}

func ListPreachers() ([]*Preacher, error) {
	resp, err := connection.query("SELECT pid, name FROM preachers")
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	var preachers []*Preacher
	for resp.Next() {
		p := new(Preacher)
		err = resp.Scan(&p.pID, &p.name)
		if err != nil {
			return nil, err
		}
		preachers = append(preachers, p)
	}
	return preachers, nil
}

func (p *Preacher) String() string {
	return p.name
}

func (p *Preacher) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PID  int    `json:"id"`
		Name string `json:"name"`
	}{
		p.pID,
		p.name,
	})
}
