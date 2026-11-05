package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomJson(c *gin.Context, status int, obj interface{}) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(status)

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Writer.Write(buffer.Bytes())
}

func JSON(c *gin.Context, err error, data interface{}, responStatus ...int) {
	if len(responStatus) > 0 && responStatus[0] != 500 {
		if err != nil {
			c.JSON(responStatus[0], gin.H{"message": err.Error()})
			return
		}
		c.JSON(responStatus[0], data)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "uuups..."})
		return
	}
	c.JSON(http.StatusOK, data)
}