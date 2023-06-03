package server

func NewStorage() *Storage {
	return &Storage{Localy: make(map[string]bool)}
}

func (s *Storage) Add(address string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.Localy[address] = true
}

func (s *Storage) Remove(address string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.Localy, address)
}

func (s *Storage) In(address string) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if val, ok := s.Localy[address]; ok {
		if val == true {
			return true
		}
	}
	return false
}
