package utils

// Based on https://flaviocopes.com/golang-data-structure-set/

// IntSet the set of Items
type IntSet struct {
	items map[int]bool
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *IntSet) Add(t int) {
	if s.items == nil {
		s.items = make(map[int]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
}

// Clear removes all elements from the Set
func (s *IntSet) Clear() {
	s.items = make(map[int]bool)
}

// Delete removes the Item from the Set and returns Has(Item)
func (s *IntSet) Delete(item int) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the Item
func (s *IntSet) Has(item int) bool {
	_, ok := s.items[item]
	return ok
}

// Items returns the Item(s) stored
func (s *IntSet) Items() []int {
	items := []int{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *IntSet) Size() int {
	return len(s.items)
}
