package main

var exists = struct{}{}

type set struct {
    elements map[string]struct{}
}

func newSet() *set {
    s := &set{}
    s.elements = make(map[string]struct{})
    return s
}

func (s *set) add(value string) {
    s.elements[value] = exists
}

func (s *set) remove(value string) {
    delete(s.elements, value)
}

func (s *set) contains(value string) bool {
    _, c := s.elements[value]
    return c
}

func (s *set) isEmpty() bool {
	return len(s.elements) == 0
}
