package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = setup()

func main() {
	router.POST("/coupons/issued", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())

			return
		}

		fmt.Printf("----------------------------------------")
		fmt.Printf("HERE IS THE OBJECT JSON: %s\n", jsonData)
		fmt.Printf("----------------------------------------")

		c.String(http.StatusOK, "OK")
	})
	router.Run()
}

func setup() *gin.Engine {
	r := gin.Default()

	bytes, err := ioutil.ReadFile("/etc/liferay/lxc/dxp-metadata/com.liferay.lxc.dxp.domains")

	if err != nil {
		log.Printf("Could not load origins config: %s\n", err)

		return r
	}

	origins := strings.Fields(string(bytes))

	var originsWithProtocol []string
	for _, x := range origins {
		originsWithProtocol = append(originsWithProtocol, "https://"+x)
	}

	log.Printf("Loading origins %s", originsWithProtocol)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     originsWithProtocol,
		AllowMethods:     []string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return r
}
