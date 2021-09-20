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

func (d *Depot) Get(id string, t string) (interface{}, error) {
	return nil, nil
}

func (d *Depot) Put(id string, t string, data interface{}) error {
	return nil
}

func (d *Depot) Find(q []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (d *Depot) Open() error {
	return d.adapter.Open()
}

func (d *Depot) Close() error {
	d.adapter.Close()
	return nil
}

func (d *Depot) NewSection(name string) *section {

	return nil
}

type section struct {
	name     string
	sections map[string]section
}

func (s *section) get(id string) (interface{}, error) {
	return nil, nil
}

func (s *section) put(id string, object interface{}) error {
	return nil
}

func (s *section) find(parameters []interface{}) ([]interface{}, error) {
	return nil, nil
}

func NewDepot(l Logisticator) *Depot {
	d := Depot{
		adapter: l,
	}
	return &d
}
