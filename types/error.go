package types

type ValidationError struct {
	Error error `json:"error"`
}

type CommonError struct {
	Error string `json:"error"`
}
