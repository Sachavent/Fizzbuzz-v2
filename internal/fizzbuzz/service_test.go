package fizzbuzz

import "testing"

type fakeRepository struct {
	incremented []GetResultQuery
	mostParams  GetResultQuery
	mostCount   int
}

func (repository *fakeRepository) IncrementRequest(query GetResultQuery) {
	repository.incremented = append(repository.incremented, query)
}

func (repository *fakeRepository) GetMostFrequentRequest() (GetResultQuery, int) {
	return repository.mostParams, repository.mostCount
}

func TestIsMultiple(t *testing.T) {
	if !isMultiple(4, 2) {
		t.Fatalf("expected true for 4 %% 2 == 0")
	}
	if isMultiple(4, 3) {
		t.Fatalf("expected false for 4 %% 3 != 0")
	}
}

func TestApplyRules(t *testing.T) {
	if got := applyRules(12, 4, 5, "fizz", "buzz"); got != "fizz" {
		t.Fatalf("expected fizz, got %q", got)
	}
	if got := applyRules(15, 4, 5, "fizz", "buzz"); got != "buzz" {
		t.Fatalf("expected buzz, got %q", got)
	}
	if got := applyRules(15, 3, 5, "fizz", "buzz"); got != "fizzbuzz" {
		t.Fatalf("expected fizzbuzz, got %q", got)
	}
	if got := applyRules(7, 3, 5, "fizz", "buzz"); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestGetResult(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	params := GetResultQuery{
		Int1:  2,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	got := service.GetResult(params)
	want := "1,fizz,3,fizz,buzz,fizz,7,fizz,9,fizzbuzz"
	if got != want {
		t.Fatalf("unexpected result: %q", got)
	}

	if len(repository.incremented) != 1 {
		t.Fatalf("expected 1 increment, got %d", len(repository.incremented))
	}
	if repository.incremented[0] != params {
		t.Fatalf("unexpected increment params: %+v", repository.incremented[0])
	}
}

func TestGetMostFrequentRequest(t *testing.T) {
	repo := &fakeRepository{
		mostParams: GetResultQuery{Int1: 2, Int2: 3, Limit: 5, Str1: "a", Str2: "b"},
		mostCount:  7,
	}
	service := NewService(repo)

	params, count := service.GetMostFrequentRequest()
	if params != repo.mostParams || count != repo.mostCount {
		t.Fatalf("unexpected most frequent result")
	}
}
