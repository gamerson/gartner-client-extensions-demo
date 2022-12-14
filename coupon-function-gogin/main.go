package main

import (
	"encoding/json"
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

func CouponIssued(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")

	fmt.Println("Authorization: " + authorization)

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var couponObject map[string]interface{}
	if err := json.Unmarshal(jsonData, &couponObject); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	objectEntry := couponObject["objectEntry"].(map[string]interface{})
	values := objectEntry["values"].(map[string]interface{})
	issued := values["issued"].(bool)

	var modifiedDate string
	if value, isMapContainsKey := objectEntry["modifiedDate"].(string); isMapContainsKey {
		modifiedDate = value
	}

	createDate := objectEntry["createDate"].(string)
	code := values["code"].(string)
	statusByUserName := objectEntry["statusByUserName"].(string)

	var status = "not issued"
	if issued {
		status = "issued"
	}

	var updatedDate = modifiedDate
	if updatedDate == "" {
		updatedDate = createDate
	}

	msg := fmt.Sprintf("The status coupon '%s' changed to '%s' by '%s' at '%s'", code, status, statusByUserName, updatedDate)

	fmt.Println(msg)

	c.String(http.StatusOK, "OK")
}

func WorkflowAction(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")

	fmt.Println("Authorization: " + authorization)

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	msg := fmt.Sprintf("JSON: %s", jsonData)

	fmt.Println(msg)

	c.String(http.StatusOK, "OK")
}

func Ready(c *gin.Context) {
	c.String(http.StatusOK, "READY")
}

func main() {
	router.GET("/", Ready)
	router.POST("/coupons/issued", CouponIssued)
	router.POST("/workflow/action", WorkflowAction)

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
