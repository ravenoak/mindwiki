package webui

type SiteLinks struct {
	Index         string
	PageIndex     string
	ProjectIndex  string
	TagIndex      string
	WebLinkIndex  string
	PageSearch    string
	ProjectSearch string
	TagSearch     string
	WebLinkSearch string
}

type SiteData struct {
	Links *SiteLinks
}

type SitePage struct {
	*SiteData
	PageTitle string
}
