package nodes

type ProjectStatus uint8

const (
	ProjectStatusUndefined ProjectStatus = iota
	ProjectStatusNotStarted
	ProjectStatusInProgress
	ProjectStatusBlocked
	ProjectStatusCompleted
)

func (s ProjectStatus) String() string {
	switch s {
	case ProjectStatusUndefined:
		return "Undefined"
	case ProjectStatusNotStarted:
		return "Not Started"
	case ProjectStatusInProgress:
		return "In Progress"
	case ProjectStatusBlocked:
		return "Blocked"
	case ProjectStatusCompleted:
		return "Completed"
	}
	return "Unknown"
}

type Project struct {
	Name        string
	Description string
	Slug       string
	Status     ProjectStatus
	StatusLine string
	Pages      []*Page
	Tags       []*Tag
}
