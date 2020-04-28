package apis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	rd "myframe/assets"
	db "myframe/database"
	"myframe/models"
)

func OrdersIndex(c *gin.Context) {
	conn := rd.Redis
	_, err := conn.Do("SET", "name", "go conf set one")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Do("expire", "name", 30)
	if err != nil {
		log.Fatalln(err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, name)
}

func OrdersOne(c *gin.Context) {
	//var mdl models.Orders
	//id, err := parseParamID(c)
	//if handleError(c, err) {
	//	return
	//}
	jsonData(c, 1)
	//mdl.Id = id
	//data, err := mdl.One()
	//if handleError(c, err) {
	//	return
	//}
	//jsonData(c, data)
}

func OrdersAll(c *gin.Context) {
	data.InitBase(c)
	var orders = make([]models.Orders, 0)
	err := db.SqlDB.Find(&orders).Error
	if err != nil {
		jsonError(c, err.Error())
	}
	data.Total = len(orders)
	var list = orders[data.Offset:data.Offset_size]
	jsonPagination(c, list)
}

//订单详情
func OrdersDetails(c *gin.Context) {
	//获取传入参数
	data.InitBase(c)
	//格式化
	res := map[string]interface{}{
		"orders": map[string]string{
			"orders_sn":   "",
			"create_time": "",
			"user_name":   "",
		},
		"status": map[string]interface{}{
			"finish": []map[string]string{
				{
					"label": "",
					"time": "",
					"status": "",
				},
			},
			"step":      "",
			"label":     "",
			"time":      "",
			"zong_time": "",
		},
		"mater": map[string]interface{}{
			"status": "1",
			"message": map[string]interface{}{
				"ex_most": map[string]string{
					"time":    "",
					"ftime":   "",
					"context": "",
				},
				"ex_all": []map[string]string{
					{
						"time":    "",
						"ftime":   "",
						"context": "",
					},
				},
			},
		},
		"site": map[string]string{
			"name":      "",
			"province":  "",
			"city":      "",
			"district":  "",
			"address":   "",
			"telephone": "",
			"mobile":    "",
		},
		"bill": map[string]string{
			"status":       "",
			"type":         "",
			"title":        "",
			"content":      "",
			"code":         "",
			"address":      "",
			"phone":        "",
			"bank_name":    "",
			"bank_account": "",
		},
		"pay": map[string]string{
			"status":         "",
			"type":           "",
			"type_name":      "",
			"time":           "",
			"code":           "",
			"payment_day":    "",
			"payment_amount": "",
		},
		"express": map[string]string{
			"type":   "",
			"amount": "",
			"number": "",
		},
		"remark": "",
		"goods": []map[string]string{
			{
				"id":              "",
				"goods_id":        "",
				"goods_sn":        "",
				"goods_name":      "",
				"goods_number":    "",
				"market_price":    "",
				"goods_price":     "",
				"promotion_type":  "",
				"goods_weight":    "",
				"goods_thumb":     "",
				"goods_attribute": "",
				"create_time":     "",
			},
		},
		"goods_amount":   "",
		"coupon":         "",
		"rebate":         "",
		"dedu":           "",
		"express_amount": "",
		"payment_amount": "",
		"goods_weight":   "",
		"comment":        "",
		"orders_status":  "",
		"isagree":        "",
		"type":           "",
		"state":          "",
		"pay_tail": map[string]string{
			"status":         "",
			"type":           "",
			"payment_amount": "",
			"payment_time":   "",
		},
		"pay_type":      "",
		"net_amount":    "",
		"events_status": "",
	}

	//获取某个订单
	var orders models.Orders
	err:=db.SqlDB.Where("orders_sn=?","3194201807021112347924407").First(&orders).Error
	if err!=nil{
		jsonError(c,err.Error())
	}
	//订单信息
	res["orders_sn"]=orders.OrdersSn
	res["create_time"]=strconv.Itoa(orders.CreateTime)
	res["user_name"]=orders.Consignee
	//订单下单流程信息
	//res["finish"]=

	//string到int
	//int,err:=strconv.Atoi(string)
	//#string到int64
	//int64, err := strconv.ParseInt(string, 10, 64)
	//#int到string
	//string:=strconv.Itoa(int)
	//#int64到string
	//string:=strconv.FormatInt(int64,10)


	//de:= struct {
	//	Name  int `json:"name-field,string,1" form:"name" val:"1"`
	//	Age  int `json:"age,string" form:"123" db:"account"`
	//	Sex   struct  {
	//		Jj string `json:"jj" form:"jj"`
	//		Ff string `json:"ff" form:"ff"`
	//		Ll struct {
	//			Gg string `json:"gg" form:"gg" `
	//			//kk string=123
	//		}`json:"ll"`
	//	}`json:"sex"`
	//}{}

	//var f interface{}
	//err := json.Unmarshal(b, &f)
	//f := map[string]interface{}{
	//	"name": "Wednesday",
	//	"age":  "6",
	//	"parents": map[string]interface{}{
	//		"sasa":  "Gomez",
	//		"sasas": "Morticia",
	//		"kankan": []map[string]interface{}{
	//			{"mama": "sasas"},
	//			{"mama": "sasas"},
	//			{"mama": "sasas"},
	//		},
	//	},
	//}
	//de.Sex.Ff="sasas"
	//de.Sex.Jj="sasas"
	//de.Name="name"
	//de.Age="name"
	//de.Sex.Ll.Gg="fafafasa"
	//c.JSON(http.StatusOK, f)
	jsonData(c, res)
	//de.sex= struct {
	//	orders string
	//}{"kanksansasa"}

	//ms:=map[string]interface{"kls":"sas",}

	//s:=[]struct{
	//	name string
	//	age string
	//}{
	//	{"user1","12"},
	//	{"user2","12"},
	//}
	//
	//var a =struct{x int}{}
	//a.x=123
	//
	//
	//var hr=gin.H{"name":de,"ji":a,"kkk":"123","ggg":s}
	////snapshot := make([]interface{}, 0)
	////snapshot[1]=lis
	//c.JSON(http.StatusOK, hr)
	//jsonData(c,hr)
}
