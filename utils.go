package collection_endpoint

type keySet struct {
	set map[string]struct{}
}

func (s *keySet) has(key string) bool {
	_, ok := s.set[key]
	return ok
}

func newKeySet[K ~string](s []K) keySet {
	var m = make(map[string]struct{})
	for _, v := range s {
		if v == "" {
			continue
		}
		m[string(v)] = struct{}{}
	}
	return keySet{set: m}
}