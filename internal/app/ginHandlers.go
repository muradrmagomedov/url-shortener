package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// shortURL() принимает метод POST и возвращает короткую ссылку

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	gin.ResponseWriter
	data *responseData
}

func (r *loggingResponseWriter) Write(body []byte) (int, error) {
	size, err := r.ResponseWriter.Write(body)
	r.data.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.data.status = statusCode
	log.Info("status ", r.data.status)
}

func ginGetURL(c *gin.Context) {
	id := c.Param("id")
	URL := urlDatabase[id]
	c.Header("Location", URL)
	c.Redirect(http.StatusTemporaryRedirect, URL)
}

func ginShortURL(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	URL := string(body)
	shortenedURLId := shortenerGenerator(shortenerLength)
	urlDatabase[shortenedURLId] = URL
	c.Header("Content-Type", "text/plain")
	if shortURLAddr == "" {
		shortURLAddr = "http://" + host + ":" + port
	}
	c.Data(http.StatusCreated, "text/plain", []byte(shortURLAddr+"/"+shortenedURLId))
}

func ginJSONShorter(c *gin.Context) {
	type LongURL struct {
		URL string `json:"url"`
	}
	type ShortURL struct {
		Result string `json:"result"`
	}

	var longURL LongURL
	var shortUrl ShortURL

	var test struct{}

	json.Unmarshal([]byte("test"), &test)

	err := c.ShouldBindBodyWithJSON(&longURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortenedURLId := shortenerGenerator(shortenerLength)
	urlDatabase[shortenedURLId] = longURL.URL
	if shortURLAddr == "" {
		shortURLAddr = "http://" + host + ":" + port

		shortUrl.Result = shortURLAddr + "/" + shortenedURLId
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusCreated, shortUrl)
	}
}

func GinLogger(h func(*gin.Context)) func(*gin.Context) {
	logFn := func(c *gin.Context) {
		start := time.Now()
		uri := c.Request.RequestURI
		duration := time.Since(start)
		method := c.Request.Method
		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := &loggingResponseWriter{
			ResponseWriter: c.Writer,
			data:           responseData,
		}
		c.Writer = lw
		h(c)
		log.Info("uri ", uri, " method ", method, " duration ", duration, " size ", lw.data.size, " status ", lw.data.status)
	}
	return logFn
}
