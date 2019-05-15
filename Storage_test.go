package main

import (
	"testing"
)

func contains(xs []RequestHistory, e RequestHistory) bool {
	for _, x := range xs {
		if x.From == e.From && x.Data == e.Data && x.Method == e.Method {
			return true
		}
	}
	return false
}

func TestStorageIsAbleToRetreiveAfterAdd(t *testing.T) {
	testsCase := []RequestHistory{
		RequestHistory{Method: "GET", From: "testa", Data: "{}"},
	}
	for _, tc := range testsCase {
		storage := StorageImpl{Store: []RequestHistory{}}
		storage.Add(tc.Method, tc.From, tc.Data)
		history := storage.GetAll()
		if !contains(history, tc) {
			t.Errorf("Storage unable to store in testcase %s", tc.From)
		}
	}
}

func TestStorageIsAbleToBeClear(t *testing.T) {
	testsCase := []RequestHistory{
		RequestHistory{Method: "GET", From: "testa", Data: "{}"},
	}
	for _, tc := range testsCase {
		storage := StorageImpl{Store: []RequestHistory{}}
		storage.Add(tc.Method, tc.From, tc.Data)
		storage.Clear()
		history := storage.GetAll()
		if len(history) != 0 {
			t.Errorf("Storage fail to Clear")
		}
	}
}
