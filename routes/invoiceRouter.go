package routes

import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)


func InvoiceRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.Get("/invoices", controller.GetInvoices())
	incomingRoutes.Get("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}