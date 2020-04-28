package routers

import (
	"github.com/gin-gonic/gin"
	."myframe/apis"
)

var Router  = gin.Default()

func init() {
	Router = gin.Default()
	Router.Static("/public", "./public")
	person:=Router.Group("person")
	{
		person.GET("/", PersonIndex)
		person.POST("/post",  PersonPost)
		person.GET("/all",  PersonAll)
		person.GET("/one",  PersonOne)
		person.PUT("/put",  PersonPut)
		person.DELETE("/delete",  PersonDelete)
	}
	orders:=Router.Group("orders")
	{
		orders.GET("/",OrdersIndex)
		//orders.POST("/post", Post)
		orders.GET("/all", OrdersAll)
		orders.GET("/one", OrdersOne)
		//orders.PUT("/put", Put)
		//orders.DELETE("/delete", Delete)
		orders.GET("/details", OrdersDetails)
	}
	pay:=Router.Group("pay")
	{
		pay.GET("/wxpay",PayWx)
		pay.Any("/alipay",PayAli)
		//pay.GET("/return",PayReturn)
		//pay.GET("/notify",PayNotify)
	}
}