package models
import (
	"sync"
	"strings"
	"github.com/spf13/cast"
	"github.com/beego/beego/v2/core/logs"
)

type KeyValueData struct {
    mu       sync.Mutex
    store map[string]string
}

var KeyValueStore = initializeDatastore()

type dataOperations interface{
	GetKey() string
	SetKey() string
	SearchPrefix() string
	SearchSuffix() string
}

func initializeDatastore() KeyValueData {
	data := KeyValueData{}
	data.store = make(map[string]string)
	data.store["sample"] = "sample Value"
	return data
}


func (s *KeyValueData) GetKey(key string) string {
	log := logs.NewLogger()
	log.Debug("Getting Key from key stroe ", key)
	s.mu.Lock()
	defer s.mu.Unlock()
	if val, ok := s.store[key]; ok {
		return val
	}
	return ""
}

func (s *KeyValueData) SetKey(key, value string) bool {
	//log := logs.NewLogger()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = cast.ToString(value)
	return true
}


func (s *KeyValueData) SearchSuffix(suffix string) []string {
	//log := logs.NewLogger()
	var searchResult []string
	s.mu.Lock()
	defer s.mu.Unlock()
	for storeKey := range s.store {
		if strings.HasSuffix(storeKey, suffix) {
			searchResult = append(searchResult, storeKey)
		}
	}
	return searchResult
}

func (s *KeyValueData) SearchPrefix(prefix string) []string {
	//log := logs.NewLogger()
	var searchResult []string
	s.mu.Lock()
	defer s.mu.Unlock()
	for storeKey := range s.store {
		if strings.HasPrefix(storeKey, prefix) {
			searchResult = append(searchResult, storeKey)
		}
	}
	return searchResult
}