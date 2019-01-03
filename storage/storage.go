package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const fileName string = "data.json"

type Tuple struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Storage interface {
	Set(tuple Tuple) Tuple
	Get(key string) (Tuple, bool)
	GetAll() []Tuple
	Del(key string)
}

type InmemoryStorage struct {
	storage map[string]string
}

func NewInmemoryStorage() InmemoryStorage {
	return InmemoryStorage{make(map[string]string)}
}

func (is InmemoryStorage) Set(tuple Tuple) Tuple {
	is.storage[tuple.Key] = tuple.Value
	return tuple
}

func (is InmemoryStorage) Get(key string) (Tuple, bool) {
	value, ok := is.storage[key]
	if ok {
		return Tuple{key, value}, true
	} else {
		return Tuple{}, false //TODO fix it!!!
	}
}

func (is InmemoryStorage) GetAll() []Tuple {
	tuples := make([]Tuple, len(is.storage))
	for key, value := range is.storage {
		tuples = append(tuples, Tuple{key, value})
	}
	return tuples
}

func (is InmemoryStorage) Del(key string) {
	delete(is.storage, key)
}

type DiskStorage struct {
}

func NewDiskStorage() DiskStorage {
	_, err := ioutil.ReadFile(fileName)
	if err != nil {
		file, _ := os.Create(fileName)
		file.Close()
	}
	return DiskStorage{}
}

func (ds DiskStorage) Set(tuple Tuple) Tuple {
	storage := readMapFromFile()
	storage[tuple.Key] = tuple.Value
	writeMapToFile(storage)
	return tuple
}

func (ds DiskStorage) Get(key string) (Tuple, bool) {
	m := readMapFromFile()
	value, ok := m[key]
	if ok {
		return Tuple{key, value}, true
	} else {
		return Tuple{}, false //TODO fix it!!!
	}
}

func (ds DiskStorage) GetAll() []Tuple {
	m := readMapFromFile()
	tuples := make([]Tuple, len(m))
	for key, value := range m {
		tuples = append(tuples, Tuple{key, value})
	}
	return tuples
}

func (ds DiskStorage) Del(key string) {
	m := readMapFromFile()
	delete(m, key)
	writeMapToFile(m)
}

func writeMapToFile(m map[string]string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	jsonMap, _ := json.Marshal(m)
	file.WriteString(string(jsonMap))
}

func readMapFromFile() map[string]string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return make(map[string]string) //TODO fix it
	}
	jsonMap := make(map[string]string)
	if len(data) == 0 {
		return jsonMap
	}
	errr := json.Unmarshal([]byte(string(data)), &jsonMap)
	if errr != nil {
		return jsonMap
	}
	return jsonMap
}
