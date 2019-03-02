package main

import (
	"testing"
)

func simpleArrayEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w := b[k]; v != w {
			return false
		}
	}

	return true
}

func simpleDeepEqual(a map[string][]string, b map[string][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || !simpleArrayEqual(v, w) {
			return false
		}
	}

	return true
}

func contains(xs []RequestHistory, e RequestHistory) bool {
	for _, x := range xs {
		if x.From == e.From && simpleDeepEqual(x.Data, e.Data) {
			return true
		}
	}
	return false
}

func TestStorageIsAbleToRetreiveAfterAdd(t *testing.T) {
	testsCase := []RequestHistory{
		RequestHistory{From: "testa", Data: map[string][]string{}},
	}
	for _, tc := range testsCase {
		storage := StorageImpl{Store: []RequestHistory{}}
		storage.Add(tc.From, tc.Data)
		history := storage.GetAll()
		if !contains(history, tc) {
			t.Errorf("Storage unable to store in testcase %s", tc.From)
		}
	}
}

func TestStorageIsAbleToBeClear(t *testing.T) {
	testsCase := []RequestHistory{
		RequestHistory{From: "testa", Data: map[string][]string{}},
	}
	for _, tc := range testsCase {
		storage := StorageImpl{Store: []RequestHistory{}}
		storage.Add(tc.From, tc.Data)
		storage.Clear()
		history := storage.GetAll()
		if len(history) != 0 {
			t.Errorf("Storage fail to Clear")
		}
	}
}
