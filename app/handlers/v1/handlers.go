package v1

const (
	BaseRoute = "/api/v1"
)

type errorRes struct {
	Error    string `json:"error"`
	Code     int    `json:"code"`
	ErrorDis string `json:"error_description"`
}

type basicResponse struct {
	Message string `json:"message"`
}
