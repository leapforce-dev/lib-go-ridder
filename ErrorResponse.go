package ridder

// ErrorResponse stores general Ridder API error response
//
type ErrorResponse struct {
	Error      string `json:"error"`
	StackTrace string `json:"stackTrace"`
}
