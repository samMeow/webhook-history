package main

const StorageSize int = 20

type RequestHistory struct {
	Method string
	From   string
	Data   string
}

type Storage interface {
	Add(method string, from string, data string)
	Clear()
	GetAll() []RequestHistory
}

type StorageImpl struct {
	Store []RequestHistory
}

// Add data to storage
func (s *StorageImpl) Add(method string, from string, data string) {
	datum := RequestHistory{
		Method: method,
		From:   from,
		Data:   data,
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
