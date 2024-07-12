package omap

type Omap[K comparable, T any] struct {
	m map[K]*omapItem[K, T]

	first *omapItem[K, T]
	last  *omapItem[K, T]
}

type omapItem[K comparable, T any] struct {
	key        K
	data       T
	prev, next *omapItem[K, T]
}

// New - initiate new Omap entity
func New[K comparable, T any]() *Omap[K, T] {
	return &Omap[K, T]{
		m: make(map[K]*omapItem[K, T]),
	}
}

// WithCap - set capacity for existing Omap entity
func (om *Omap[K, T]) WithCap(c int) *Omap[K, T] {
	om.m = make(map[K]*omapItem[K, T], c)
	return om
}

// Len - return len for existing Omap entity
func (om *Omap[K, T]) Len() int {
	return len(om.m)
}

// Get - get data by key in Comma Ok
func (om *Omap[K, T]) Get(key K) (T, bool) {
	v, ok := om.m[key]
	if !ok {
		return omapItem[K, T]{}.data, ok
	}
	return v.data, ok
}

// Set - set data with key
// if key already exists, deleting the key and set new object to the end of the list
func (om *Omap[K, T]) Set(key K, value T) {
	om.Delete(key)

	oi := &omapItem[K, T]{
		key:  key,
		data: value,
	}
	if om.first == nil {
		om.first = oi
	}
	if om.last == nil {
		om.last = oi
	} else {
		om.last.next = oi
		oi.prev = om.last
		om.last = oi
	}

	om.m[key] = oi
}

// Delete - delete data by key, returning bool
// true - if deletion was successful
func (om *Omap[K, T]) Delete(key K) bool {
	v, ok := om.m[key]
	if !ok {
		return false
	}

	if v.next != nil {
		if v == om.first {
			om.first = v.next
		}
		v.next.prev = v.prev
	}
	if v.prev != nil {
		if v == om.last {
			om.last = v.prev
		}
		v.prev.next = v.next
	}

	if v == om.first {
		om.first = nil
	}
	if v == om.last {
		om.last = nil
	}

	delete(om.m, key)

	return true
}

// Replace - replace data with key, returning bool
// have no change the existing order
// true - if replace was successful
// false - if no replace and set
func (om *Omap[K, T]) Replace(key K, value T) bool {
	v, ok := om.m[key]
	if !ok {
		return false
	}

	v.data = value
	return true
}

// Iter - ordered forward iter
// accepting key-value func
func (om *Omap[K, T]) Iter(f func(key K, value T)) {
	if om.first == nil {
		return
	}
	for e := om.first; e != nil; e = e.next {
		f(e.key, e.data)
	}
}

// IterBack - ordered backward iter
// accepting key-value func
func (om *Omap[K, T]) IterBack(f func(key K, value T)) {
	if om.last == nil {
		return
	}
	for e := om.last; e != nil; e = e.prev {
		f(e.key, e.data)
	}
}
