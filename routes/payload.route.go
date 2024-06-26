package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

type PayloadRoutes struct {
	payloadController controllers.PayloadController
}

func NewRoutePayload(payloadController controllers.PayloadController) PayloadRoutes {
	return PayloadRoutes{payloadController}
}

func (pr *PayloadRoutes) PayloadRoute(rg *gin.RouterGroup) {
	router := rg.Group("payloads")
	router.POST("/", pr.payloadController.CreatePayload)
	router.GET("/", pr.payloadController.GetUndoPayloads)
	router.PATCH("/:payloadId", pr.payloadController.UpdatePayload)
}
