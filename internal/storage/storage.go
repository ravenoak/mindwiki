package storage

type Logisticator interface {
	Delete(set, id []byte) (interface{}, error)
	Insert(set, id []byte, object interface{}) error
	Query(parameters []interface{}) ([]interface{}, error)

	Open() error
	Close()
}

type Depot struct {
	adapter  Logisticator
	sections map[string]section
}

func (d *Depot) Get(id string) (interface{}, error) {
	return nil, nil
}

func (d *Depot) Put(id string, object interface{}) error {
	return nil
}

func (d *Depot) Find(parameters []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (d *Depot) Open() error {
	return d.adapter.Open()
}

func (d *Depot) Close() {
	d.adapter.Close()
}

func (d *Depot) NewSection(name string) *section {

	return nil
}

type section struct {
	name     string
	sections map[string]section
}

func (s *section) Get(id string) (interface{}, error) {
	return nil, nil
}

func (s *section) Put(id string, object interface{}) error {
	return nil
}

func (s *section) Find(parameters []interface{}) ([]interface{}, error) {
	return nil, nil
}

func NewDepot(l Logisticator) *Depot {
	d := Depot{
		adapter: l,
	}
	return &d
}
