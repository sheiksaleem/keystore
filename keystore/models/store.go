package models

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/spf13/cast"
	"strings"
	"sync"
	//"github.com/beego/beego/v2/core/logs"
)

type KeyValueData struct {
	mu    sync.Mutex
	store map[string]string
}

var KeyValueStore = initializeDatastore()

type dataOperations interface {
	GetKey() string
	SetKey() string
	SearchPrefix() string
	SearchSuffix() string
}

func initializeDatastore() KeyValueData {
	data := KeyValueData{}
	data.store = make(map[string]string)
	if beego.BConfig.RunMode == "dev" {
		data.store["abc-1"] = "abc-1 Value"
		data.store["abc-2"] = "abc-2 Value"
		data.store["xyz-1"] = "xyz-1 Value"
		data.store["xyz-2"] = "xyz-2 Value"
	}
	return data
}

func (s *KeyValueData) GetKey(key string) string {
	// log := logs.NewLogger()
	// log.Debug("Getting Key from key stroe ", key)
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
