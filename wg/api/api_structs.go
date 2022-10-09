package api

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrAPI = errors.New("api error")
)

type Request struct {
	Region Region
	App    App
	Method Method
	Data   url.Values
}

type ResponseError struct {
	Code    int    `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

func (e *ResponseError) Text() string {
	lines := []string{
		strconv.Itoa(e.Code) + ": " + e.Message,
	}

	if e.Field != "" || e.Value != "" {
		lines = append(lines, e.Field+" "+e.Value)
	}

	return strings.Join(lines, "\n")
}

func (e *ResponseError) GetError() error {
	return fmt.Errorf("%w %s", ErrAPI, e.Text())
}

type ResponseMeta struct {
	Conunt int `json:"count"`
}

type Response struct {
	Status string         `json:"status"`
	Error  *ResponseError `json:"error"`
	Meta   ResponseMeta   `json:"meta"`
	Data   interface{}    `json:"data"`
}
