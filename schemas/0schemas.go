// Package schemas for Request and Response Schema
package schemas

// APIResponse Default API Response
type APIResponse struct {
	OK      bool   `json:"ok"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetsResponse Default API Response for get data multiple result
type GetsResponse[T any] struct {
	APIResponse
	Data  []T `json:"data"`
	Total int `json:"total"`
}

// GetOneResponse Default API Response for get date by id
type GetOneResponse[T any] struct {
	APIResponse
	Data  T   `json:"data"`
	Total int `json:"total"`
}
