package pkg

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

const (
	// HTTPClientTimeout is the timeout for HTTP requests.
	HTTPClientTimeout = 30 * time.Second

	// HTTPClientUserAgent is the user agent for HTTP requests.
	HTTPClientUserAgent = "Backend Vaccine HP for JHCIS API Service"
)

// FastHTTPGet sends a GET request to the specified URL with the given headers.
func FastHTTPGet(url string, headers map[string]string) ([]byte, error) {
	return FastHTTPRequest(url, fasthttp.MethodGet, nil, headers)
}

// FastHTTPPost sends a POST request to the specified URL with the given body and headers.
func FastHTTPPost(url string, body []byte, headers map[string]string) ([]byte, error) {
	return FastHTTPRequest(url, fasthttp.MethodPost, body, headers)
}

// FastHTTPRequest makes a fast HTTP request to the specified URL using the given method, body, and headers.
//
// Parameters:
//
// - url: The URL to send the HTTP request to.
//
// - method: The HTTP method to use for the request.
//
// - body: The body of the HTTP request.
//
// - headers: The headers to include in the HTTP request.
//
// Return:
//
// - []byte: The response body of the HTTP request.
func FastHTTPRequest(url string, method string, body []byte, headers map[string]string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.SetTimeout(HTTPClientTimeout)
	req.Header.SetMethod(method)
	req.Header.Set("User-Agent", HTTPClientUserAgent)
	if method != fasthttp.MethodGet {
		req.SetBody(body)
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		log.Warn().Err(err).Int("status", resp.StatusCode()).Msg("HTTP request failed")
		return resp.Body(), err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		log.Warn().Err(err).Int("status", resp.StatusCode()).Msg("HTTP request failed")
		return resp.Body(), fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode())
	}

	contentEncoding := resp.Header.Peek("Content-Encoding")
	if len(contentEncoding) > 1 {
		return resp.Body(), nil
	}

	return resp.Body(), nil
}
