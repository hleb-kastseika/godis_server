package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const fileName string = "data.json"

type Tuple struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Common Storage interface
type Storage interface {
	Set(tuple Tuple) Tuple
	Get(key string) (Tuple, bool)
	GetAll() []Tuple
	Del(key string)
	FindKeys(keyPattern string) ([]string, bool)
}

//InmemoryStorage - Storage interface implementation which works with inmemory map
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
	return get(is.storage, key)
}

func (is InmemoryStorage) GetAll() []Tuple {
	return getAll(is.storage)
}

func (is InmemoryStorage) Del(key string) {
	delete(is.storage, key)
}

func (is InmemoryStorage) FindKeys(keyPattern string) ([]string, bool) {
	return findKeys(keyPattern, is.storage)
}

//DiskStorage - Storage interface implementation which works with map stored in file on the disk
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
	return get(m, key)
}

func (ds DiskStorage) GetAll() []Tuple {
	m := readMapFromFile()
	return getAll(m)
}

func (ds DiskStorage) Del(key string) {
	m := readMapFromFile()
	delete(m, key)
	writeMapToFile(m)
}

func (ds DiskStorage) FindKeys(keyPattern string) ([]string, bool) {
	m := readMapFromFile()
	return findKeys(keyPattern, m)
}

func get(m map[string]string, key string) (Tuple, bool) {
	value, ok := m[key]
	if ok {
		return Tuple{key, value}, true
	} else {
		return Tuple{}, false //TODO fix it!!!
	}
}

func getAll(m map[string]string) []Tuple {
	tuples := make([]Tuple, len(m))
	for key, value := range m {
		tuples = append(tuples, Tuple{key, value})
	}
	return tuples
}

func findKeys(keyPattern string, m map[string]string) ([]string, bool) {
	keys := make([]string, len(m))
	keyPattern = strings.Replace(keyPattern, "*", "(.*)", -1)
	for key := range m {
		match, _ := regexp.MatchString(keyPattern, key)
		if match {
			keys = append(keys, key)
		}
	}
	return keys, len(keys) > 0
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
