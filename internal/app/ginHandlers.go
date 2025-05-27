package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// shortURL() принимает метод POST и возвращает короткую ссылку
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
	c.Data(http.StatusCreated, "text/plain", []byte("http://"+host+":"+port+"/"+shortenedURLId))
}

func ginGetURL(c *gin.Context) {
	id := c.Param("id")
	URL := urlDatabase[id]
	c.Header("Location", URL)
	c.Redirect(http.StatusTemporaryRedirect, URL)
}
