package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGinShortURL(t *testing.T) {
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
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			ginShortURL(c)
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

func TestGinGetURL(t *testing.T) {
	mockStorage["123AbcdF"] = "https://ya.ru"
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
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			ginGetURL(c)
			result := w.Result()
			defer result.Body.Close()
			require.Equal(t, tt.want.statusCode, result.StatusCode)

		})
	}
}

func TestGinJSONShorter(t *testing.T) {
	type LongUrl struct {
		Url string `json:"url"`
	}
	type ShortURL struct {
		Result string `json:"result"`
	}
	type want struct {
		name       string
		longURL    LongUrl
		requestURL string
	}
	tests := []want{
		{name: "simple url", longURL: LongUrl{Url: "https://mail.ru"}, requestURL: "/api/shorten"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var sUrl ShortURL
			body, err := json.Marshal(test.longURL)
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, test.requestURL, bytes.NewReader(body))
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			ginJSONShorter(c)
			result := w.Result()
			require.Equal(t, http.StatusCreated, result.StatusCode)
			bodyByte, err := io.ReadAll(result.Body)
			defer result.Body.Close()
			require.NoError(t, err)
			err = json.Unmarshal(bodyByte, &sUrl)
			require.NoError(t, err)
			str := string(sUrl.Result)
			arr := strings.Split(str, "/")
			id := arr[len(arr)-1]
			wantLongURL, ok := urlDatabase[id]
			assert.True(t, ok)
			require.Equal(t, wantLongURL, test.longURL.Url)
		})
	}
}
