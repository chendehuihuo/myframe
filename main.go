package main

import (
	router "myframe/routers"
)

func run(){
		// 每5秒打印一次数据，查看自己更改配置文件后是否可以热刷新
		rou:=router.Router
		rou.Run(":8080")
		//time.Sleep(5)
}
func main() {
	for{
		run()
	}
//fmt.Println(123)
}


//type Ceshi struct {
//	Time   int  `json:"time" form:"time"`
//	Name string  `json:"name" form:"name"`
//}

//func main()  {
	//router := gin.Default()

	//SqlDB, err := gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	//SqlDB, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?parseTime=true")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	//var count int
	//err = SqlDB.Model(&name).Count(&count).Error
	//if err != nil {
	//	fmt.Println(count)
	//}
	//var name Ceshi
	//re:=SqlDB.First(&name)

	//if res != nil {
	//	fmt.Println(err)
	//}
	//var name []
	//SqlDB.Find(&name)
	//fmt.Println(name)
	//fmt.Println("sasas")
	//fmt.Println(count)

	//var c *gin.Context
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "sasa",
	//})


	//r := gin.Default()

	//r.GET("/some", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"lang": "GO语言",
		//	"tag":  "<br>",
		//}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		//c.JSON(http.StatusOK, re)
	//})

	// Listen and serve on 0.0.0.0:8080
	//r.Run(":8080")


	//router.GET("/someJSON", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"lang": "GO语言",
	//		"tag":  "<br>",
	//	}
	//
	//	// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	//	c.AsciiJSON(http.StatusOK, data)
	//})
//}