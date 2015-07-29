package intset

type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found
}

func (set *IntSet) Get(i int) bool {
	_, found := set.set[i]
	return found
}

func (set *IntSet) Remove(i int) {
	delete(set.set, i)
}
