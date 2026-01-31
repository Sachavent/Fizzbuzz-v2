package fizzbuzz

import "testing"

func TestInMemoryStorageRepository_Empty(t *testing.T) {
	repository := NewInMemoryStorageRepository().(*InMemoryStorageRepository)

	params, count := repository.GetMostFrequentRequest()
	if count != 0 {
		t.Fatalf("expected count 0, got %d", count)
	}
	if params != (GetResultQuery{}) {
		t.Fatalf("expected zero params, got %+v", params)
	}
}

func TestInMemoryStorageRepository_IncrementAndMostFrequent(t *testing.T) {
	repository := NewInMemoryStorageRepository().(*InMemoryStorageRepository)

	first := GetResultQuery{Int1: 2, Int2: 3, Limit: 10, Str1: "fizz", Str2: "buzz"}
	second := GetResultQuery{Int1: 3, Int2: 5, Limit: 20, Str1: "a", Str2: "b"}

	repository.IncrementRequest(first)
	repository.IncrementRequest(first)
	repository.IncrementRequest(second)

	params, count := repository.GetMostFrequentRequest()
	if count != 2 {
		t.Fatalf("expected count 2, got %d", count)
	}
	if params != first {
		t.Fatalf("expected params %+v, got %+v", first, params)
	}
}
