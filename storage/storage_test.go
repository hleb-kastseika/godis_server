package storage

import "testing"

func TestInmemoryStorageSet(t *testing.T) {
	storage := NewInmemoryStorage()

	testKey := "test"
	testValue := "testValue"
	storage.Set(Tuple{testKey, testValue})

	_, ok := storage.Get(testKey)
	if !ok {
		t.Errorf("Could'n find map value!")
	}
}
