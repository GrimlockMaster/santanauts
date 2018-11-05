package santanaut

import "fmt"

type Santanaut struct {
	Id        string
	Name      string
	Email     string
	Blacklist []string
	Entries   int
	Targets   []string
}

func New(id, name, email string, entries int, blacklist []string) *Santanaut {
	return &Santanaut{Id: id, Name: name, Email: email, Entries: entries, Blacklist: blacklist, Targets: make([]string, 0)}
}

func (s *Santanaut) String() string {
	//return fmt.Sprintf("Name: %s, Email: %s, Entries: %d, Blacklist: %v, Targets: %v", s.Name, s.Email, s.Entries, s.Blacklist, s.Targets)
	//return s.Name
	return fmt.Sprintf("Name: %s, Targets: %v", s.Name, s.Targets)
}

func (s Santanaut) IsValidTarget(target Santanaut) bool {
	if s.Name == target.Name {
		return false
	}

	if contains(s.Blacklist, target.Id) {
		return false
	}

	if contains(s.Targets, target.Id) {
		return false
	}

	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
