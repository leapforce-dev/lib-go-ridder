package ridder

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int32  `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}
