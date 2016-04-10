package db

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

	} else {

	}
	return nil
}
