package storage

type Page struct {
	Title string
	Body  string
	Slug  string
	Tags  []*Tag
}
