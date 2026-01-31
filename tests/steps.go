package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cucumber/godog"
)

type testContext struct {
	baseURL      string
	profileName  string
	client       *http.Client
	lastResponse *httpResponse
}

type httpResponse struct {
	statusCode int
	body       any
	rawBody    []byte
}

func InitializeScenario(scenarioContext *godog.ScenarioContext) {
	testContext := &testContext{
		baseURL: os.Getenv("BASE_URL"),
		client:  &http.Client{Timeout: 5 * time.Second},
	}
	if testContext.baseURL == "" {
		testContext.baseURL = "http://localhost:8080"
	}

	scenarioContext.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		testContext.profileName = ""
		testContext.lastResponse = nil

		return ctx, nil
	})

	scenarioContext.Step(`^I make a (\w+) request on the route "([^"]+)"$`, testContext.iMakeRequestOnRoute)
	scenarioContext.Step(`^the response status code is (\d+)$`, testContext.theResponseStatusCodeIs)
	scenarioContext.Step(`^the response body is:$`, testContext.theResponseBodyIs)
}

func (testContext *testContext) iMakeRequestOnRoute(method, route string) error {
	return testContext.doRequest(method, route, nil, nil)
}

func (testContext *testContext) doRequest(method, route string, headers map[string]string, body any) error {
	if headers == nil {
		headers = map[string]string{}
	}

	var bodyReader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to encode body: %w", err)
		}
		bodyReader = bytes.NewReader(buf)
	}

	url := route
	if !strings.HasPrefix(route, "http://") && !strings.HasPrefix(route, "https://") {
		url = strings.TrimRight(testContext.baseURL, "/") + "/" + strings.TrimLeft(route, "/")
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := testContext.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var parsed any
	if len(raw) > 0 {
		_ = json.Unmarshal(raw, &parsed)
	}
	testContext.lastResponse = &httpResponse{statusCode: resp.StatusCode, body: parsed, rawBody: raw}

	return nil
}

func (testContext *testContext) theResponseStatusCodeIs(code int) error {
	if testContext.lastResponse == nil {
		return fmt.Errorf("no response recorded")
	}
	if testContext.lastResponse.statusCode != code {
		return fmt.Errorf("expected status %d, got %d", code, testContext.lastResponse.statusCode)
	}

	return nil
}

func (testContext *testContext) theResponseBodyIs(doc *godog.DocString) error {
	if testContext.lastResponse == nil {
		return fmt.Errorf("no response recorded")
	}

	expectedRaw := strings.TrimSpace(doc.Content)
	actualRaw := strings.TrimSpace(string(testContext.lastResponse.rawBody))

	expectedJSON, err := parseJSON(doc)
	if err == nil {
		if !deepEqual(testContext.lastResponse.body, expectedJSON) {
			return fmt.Errorf(
				"response body mismatch (json)\nexpected: %s\ngot: %s",
				expectedRaw,
				actualRaw,
			)
		}

		return nil
	}

	if actualRaw != expectedRaw {
		return fmt.Errorf(
			"response body mismatch (string)\nexpected: %q\ngot: %q",
			expectedRaw,
			actualRaw,
		)
	}

	return nil
}

func parseJSON(doc *godog.DocString) (any, error) {
	if doc == nil {
		return nil, fmt.Errorf("missing docstring")
	}
	var jsonOutput any
	if err := json.Unmarshal([]byte(doc.Content), &jsonOutput); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return jsonOutput, nil
}

func deepEqual(source, toCompare any) bool {
	sourceEncoded, _ := json.Marshal(source)
	toCompareEncoded, _ := json.Marshal(toCompare)
	return bytes.Equal(sourceEncoded, toCompareEncoded)
}
