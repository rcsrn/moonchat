package server

var exists = struct{}{}

type set struct {
	elements map[string]struct{}
}

func newSet() *set {
	s := &set{}
	s.elements = make(map[string]struct{})
	return s
}

func (s *set) getElements() map[string]struct{} {
	return s.elements
}

func (s *set) add(value string) {
	s.elements[value] = exists
}

func (s *set) remove(value string) {
	delete(s.elements, value)
}

func (s *set) contains(value string) bool {
	if _, itExists := s.elements[value]; itExists {
		return true
	}
	return false
}

func (s *set) isEmpty() bool {
	return len(s.elements) == 0
}
