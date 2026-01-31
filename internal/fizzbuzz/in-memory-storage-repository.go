package fizzbuzz

import "sync"

type InMemoryStorageRepository struct {
	cache map[GetResultQuery]int
	mutex sync.RWMutex
}

func NewInMemoryStorageRepository() Repository {
	return &InMemoryStorageRepository{
		cache: make(map[GetResultQuery]int),
	}
}

func (repository *InMemoryStorageRepository) GetMostFrequentRequest() (GetResultQuery, int) {
	repository.mutex.RLock()
	defer repository.mutex.RUnlock()

	maxCount := 0
	mostFrequentQuery := GetResultQuery{}

	for query, count := range repository.cache {
		if count > maxCount {
			maxCount = count
			mostFrequentQuery = query
		}
	}

	return mostFrequentQuery, maxCount
}

func (repository *InMemoryStorageRepository) IncrementRequest(query GetResultQuery) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()

	repository.cache[query]++
}
