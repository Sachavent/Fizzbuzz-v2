package fizzbuzz

type GetResultQuery struct {
	Int1  int    `query:"int1" validate:"required,gt=0"`
	Int2  int    `query:"int2" validate:"required,gt=0"`
	Limit int    `query:"limit" validate:"required,gt=0,lte=10000"`
	Str1  string `query:"str1" validate:"required,max=10"`
	Str2  string `query:"str2" validate:"required,max=10"`
}

type GetMostFrequentRequestOutput struct {
	Parameters GetResultQuery `json:"parameters"`
	Count      int            `json:"count"`
}

type ErrorResponse struct {
	Error  string   `json:"error,omitempty"`
	Errors []string `json:"errors,omitempty"`
}
