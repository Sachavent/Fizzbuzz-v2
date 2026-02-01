package fizzbuzz

import (
	"strconv"
	"strings"
)

type ServiceInterface interface {
	GetResult(fizzBuzzParams GetResultQuery) string
	GetMostFrequentRequest() (GetResultQuery, int)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) ServiceInterface {
	return &service{repository: repository}
}

func isMultiple(iterationNumber int, multiple int) bool {
	return iterationNumber%multiple == 0
}

func applyRules(iterationNumber int, int1 int, int2 int, str1 string, str2 string) string {
	if isMultiple(iterationNumber, int1) && isMultiple(iterationNumber, int2) {
		return str1 + str2
	}

	if isMultiple(iterationNumber, int1) {
		return str1
	}

	if isMultiple(iterationNumber, int2) {
		return str2
	}

	return ""
}

func (service *service) GetResult(params GetResultQuery) string {
	var result []string

	for i := 1; i < params.Limit+1; i++ {
		element := applyRules(i, params.Int1, params.Int2, params.Str1, params.Str2)

		if element != "" {
			result = append(result, element)
		} else {
			result = append(result, strconv.Itoa(i))
		}

	}

	service.repository.IncrementRequest(params)

	return strings.Join(result, ",")
}

func (service *service) GetMostFrequentRequest() (GetResultQuery, int) {
	return service.repository.GetMostFrequentRequest()
}
