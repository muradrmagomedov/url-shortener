package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortURL(t *testing.T) {
	type want struct {
		method      string
		statusCode  int
		requestURL  string
		length      int
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{name: "test #1", want: want{requestURL: "/", statusCode: http.StatusCreated, method: http.MethodPost, length: 30, contentType: "text/plain"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			siteName := "ya.ru"
			req := httptest.NewRequest(http.MethodPost, tt.want.requestURL, strings.NewReader(siteName))
			w := httptest.NewRecorder()
			shortURL(w, req)
			result := w.Result()
			bodyByte, err := io.ReadAll(result.Body)
			defer result.Body.Close()
			body := string(bodyByte)
			require.NoError(t, err)
			require.Equal(t, tt.want.length, len(body))
			require.Equal(t, tt.want.statusCode, result.StatusCode)
			require.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			params := strings.Split(string(body), "/")
			id := params[len(params)-1]
			url, ok := urlDatabase[id]
			require.True(t, ok)
			require.Equal(t, siteName, url)
		})
	}
}

func TestGetURL(t *testing.T) {
	mockStorage["123AbcdF"] = "ya.ru"
	type want struct {
		method     string
		statusCode int
		target     string
	}
	tests := []struct {
		name string
		want want
	}{
		{name: "test #1", want: want{statusCode: http.StatusTemporaryRedirect, method: http.MethodGet, target: "/123AbcdF"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.want.target, nil)
			w := httptest.NewRecorder()
			getURL(w, req)
			result := w.Result()
			require.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}
