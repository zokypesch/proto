package main

import (
	"github.com/gin-gonic/gin"
	svc "github.com/zokypesch/kudo/service"
)

var listAction = map[string]interface{}{
	"PublishPulsa": svc.NewPublishPulsaService,
	"PublishData":  svc.NewPublishDataService,
}

func posting(c *gin.Context) {
	// var responseGetAll map[string]interface{}
	errMessage := ""
	// raw, _ := c.GetRawData()

	// json.Unmarshal([]byte(raw), &responseGetAll)

	// payload := responseGetAll["Payload"].(map[string]string)

	// req := &svc.Request{
	// 	Type:        responseGetAll["Type"].(string),
	// 	Code:        responseGetAll["Code"].(string),
	// 	TransSerial: responseGetAll["TransSerial"].(string),
	// }
	var req svc.Request
	if errBind := c.ShouldBindJSON(&req); errBind != nil {
		c.JSON(400, gin.H{
			"status":  "your request terminated",
			"message": "error params",
			"error":   errBind.Error(),
		})
	}

	// log.Println("resp", req)

	ex, exist := listAction[req.Attribute.Type]

	if !exist {
		c.JSON(400, gin.H{
			"status":  "your request terminated",
			"message": "uknow vendor",
		})
	}
	st := ex.(func() svc.Action)()
	if err := st.Do(&req); err != nil {
		errMessage = err.Error()
	}
	c.JSON(200, gin.H{
		"status": "request executed",
		"error":  errMessage,
	})
}

func main() {
	router := gin.Default()
	router.POST("/test", posting)
	router.Run()
}

// Attribute: {
// Type: PublishPulsa, PublisCode, VoucherMobileLegend
// Code: XL50
// TransSerial: xxxxCCDD
// }
// Data: Payload
