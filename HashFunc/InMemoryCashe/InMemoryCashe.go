package main

import "time"

var _ Cache = InMemoryCache{} // это трюк для проверки типа: до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn time.Duration
	storage  map[string]*CacheEntry
}

func (in InMemoryCache) Set(key string, value interface{}) {
	entry := &CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
	in.storage[key] = entry
}

func (in InMemoryCache) Get(key string) interface{} {
	entry := in.storage[key]
	if entry == nil || time.Since(entry.settledAt) > in.expireIn {
		return nil
	}
	return entry.value
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,
		storage:  make(map[string]*CacheEntry), // Изменил тип данных хранилища на map[string]*CacheEntry
	}
}
