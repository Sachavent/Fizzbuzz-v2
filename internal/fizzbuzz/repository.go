package fizzbuzz

type Repository interface {
	IncrementRequest(query GetResultQuery)
	GetMostFrequentRequest() (GetResultQuery, int)
}
