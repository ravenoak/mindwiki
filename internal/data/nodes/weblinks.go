package nodes

import (
	"net/url"
	"time"
)

type WebLink struct {
	Slug         string
	URL          url.URL
	Description  string
	Tags         []*Tag
	LastVerified time.Time
}
