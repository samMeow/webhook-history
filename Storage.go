package main

const StorageSize int = 20

type RequestHistory struct {
	From string
	Data map[string][]string
}

type Storage interface {
	Add(from string, data map[string][]string)
	Clear()
	GetAll() []RequestHistory
}

type StorageImpl struct {
	Store []RequestHistory
}

// Add data to storage
func (s *StorageImpl) Add(from string, data map[string][]string) {
	datum := RequestHistory{
		From: from,
		Data: data,
	}
	if len(s.Store) > StorageSize {
		s.Clear()
	}
	s.Store = append(s.Store, datum)
}

func (s *StorageImpl) Clear() {
	s.Store = s.Store[:0]
}

func (s *StorageImpl) GetAll() []RequestHistory {
	return s.Store
}
