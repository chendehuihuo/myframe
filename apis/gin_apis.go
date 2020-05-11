package apis

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"myframe/components"
)
var data components.Data
//
//import (
//	"errors"
//	"github.com/gin-contrib/cors"
//	"github.com/gin-gonic/gin"
//	"strconv"
//	"time"
//	"www/apiapp/models"
//)
//
//var r = gin.Default()
//错误提示
func jsonError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(200, gin.H{"code": 0, "msg": msg})
}

//不分页
func jsonData(c *gin.Context, list interface{}) {
	c.JSON(200, gin.H{"code": 1, "data": list})
}

//分页
func jsonPagination(c *gin.Context, list interface{}) {
	c.JSON(200, gin.H{"code": 1, "data": list, "total": data.Total, "page": data.Page, "page_size": data.Page_size})
}

//执行数据
func jsonSuccess(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "msg": "success"})
}

//数据转换
func conver(){

}

func tablename(str string) string{
	return "ims_"+str;
}


func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//func listpage(c *gin.Context){
//	 lis:= orders[data.Offset:data.Offset_size]
//	//var re= list[data.Page,data.Page_size]
//	jsonPagination(c,lis)
//}
//
//func handleError(c *gin.Context, err error) bool {
//	if err != nil {
//		jsonError(c, err.Error())
//		return true
//	}
//	return false
//}
//
//func parseParamID(c *gin.Context) (uint, error) {
//	id := c.Param("id")
//	parseId, err := strconv.ParseUint(id, 10, 32)
//	if err != nil {
//		return 0, errors.New("id must be an unsigned int")
//	}
//	return uint(parseId), nil
//}
//
//func enableCorsMiddleware() {
//	//TODO:: customize your own CORS
//	//https://github.com/gin-contrib/cors
//	// CORS for https://foo.com and https://github.com origins, allowing:
//	// - PUT and PATCH methods
//	// - Origin header
//	// - Credentials share
//	// - Preflight requests cached for 12 hours
//	r.Use(cors.New(cors.Config{
//		AllowOrigins:     []string{"*"}, //https://foo.com
//		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
//		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
//		ExposeHeaders:    []string{"Content-Length"},
//		AllowCredentials: false, //enable cookie
//		AllowOriginFunc: func(origin string) bool {
//			return true
//			//return origin == "https://github.com"
//		},
//		MaxAge: 12 * time.Hour, //cache options result decrease request lag
//	}))
//}
