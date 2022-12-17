package set

const (
	defaultCap = 1 << 3
)

type Set interface {
	Add(interface{}) bool
	AddAll(...interface{}) Set
	Contains(interface{}) bool
	Del(interface{}) bool
}

type set struct {
	bucket map[interface{}]struct{}
}

func NewSet() Set {
	return NewSetWithCap(defaultCap)
}

func NewSetWithCap(cap int) Set {
	return &set{bucket: make(map[interface{}]struct{}, cap)}
}

func (s *set) Add(i interface{}) bool {
	_, ok := s.bucket[i]
	s.bucket[i] = struct{}{}
	return !ok
}

func (s *set) AddAll(is ...interface{}) Set {
	for i := range is {
		s.bucket[is[i]] = struct{}{}
	}
	return s
}

func (s *set) Contains(i interface{}) bool {
	_, ok := s.bucket[i]
	return ok
}

func (s *set) Del(i interface{}) bool {
	_, ok := s.bucket[i]
	delete(s.bucket, i)
	return ok
}
