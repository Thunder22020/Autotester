package domain

// UrlRequest represents the payload for URL scanning requests.
type UrlRequest struct {
	Url      string   `json:"url" validate:"required,http_url"`
	Tests    []string `json:"tests,omitempty"`
	NeedUIUX bool     `json:"needUIUX,omitempty"`
}

// APIResponse represents a standard response structure for API endpoints.
type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// Result represents a test result.
type Result struct {
	Test    string `json:"test"`
	Result  bool   `json:"result"`
	Comment string `json:"comment,omitempty"`
}

type TestResult struct {
	Status  string   `json:"status"`
	Results []Result `json:"results"`
	Error   string   `json:"error,omitempty"`
}
