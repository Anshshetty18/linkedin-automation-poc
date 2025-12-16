package search

type Profile struct {
	Name     string
	Title    string
	Location string
	URL      string
}

type Service interface {
	FindProfiles(query string) ([]Profile, error)
}
