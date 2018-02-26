package main

import (
	//b64 "encoding/base64"
	"github.com/gin-gonic/gin"
	//"log"
)

const favicon = "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH3wUfDQE1qFaq0gAAAP1JREFUOMtjYCAB/L/EjMxWAGJVcjW7APFLIE4lWjPMACAdDcQfgPg/EPsQawAjlE4C4i9QzSDsTbTTgbQXEH9E0gzC0cRqlgXi+2iaQbgT3X8gzIQmxgjEK7Bofg3Emci2MQOxPRCHoBkQBcSf0TQfAWIDZM1SQNwEDd3fQKwMFRcH4sNQTf+g9AYgloF7EUjEAvElNBs2QRUkoomvA2IhlHQBZOgA8UUsfgwH4itI/ENALIieqGABxwPEXUjOBOFvSGxQDMhjaMaSREFOfoHmkk9A7EpKOncG4rtQzX+AuAabOlzpHZZkzYH4DhDvBnmPoGYcqU8PiI0IaQYAZQNEhD7PHnEAAAAASUVORK5CYII="

func main() {

	r := gin.Default()

	/*
		r.GET("/favicon.ico", func(c *gin.Context) {
			favicondata, _ := b64.StdEncoding.DecodeString(favicon)
			c.Writer.Header().Set("Content-Type", "image/png")
			c.Writer.Header().Set("Expires", "Tue, 31 Dec 2030 23:30:45 GMT")
			c.Writer.Header().Set("Cache-Control", "private, max-age=630720000")
			c.String(200, string(favicondata))
		})
	*/

	r.Static("/", "./public")

	r.Run(":8080")

}
