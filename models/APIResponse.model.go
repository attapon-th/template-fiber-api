package models

// APIResponse Default API Response
type APIResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
} // @name APIResponse

func NewAPIResponse(ok bool, message string) *APIResponse {
	return &APIResponse{
		OK:      ok,
		Message: message,
	}
}

func NewAPIResponseError(err error) *APIResponse {
	var message string
	if err != nil {
		message = err.Error()
	}
	return &APIResponse{
		OK:      err == nil,
		Message: message,
	}
}

func (r *APIResponse) SetResponse(ok bool, message string) *APIResponse {
	r.Message = message
	r.OK = ok
	return r
}

func (r *APIResponse) SetError(err error) *APIResponse {
	r.OK = false
	if err != nil {
		r.Message = err.Error()
	}
	return r
}

// APIResponse Default API Response
type APIResponseOne[T any] struct {
	*APIResponse
	Data *T `json:"data"`
} // @name APIResponseOne

// APIResponse Default API Response
type APIResponsePagination[T any] struct {
	*APIResponse
	Data       []T         `json:"data"`
	Pagination *Pagination `json:"pagination"`
} // @name APIResponsePagination
