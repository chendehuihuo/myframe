package components

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Data struct {
	Page int
	Page_size int
	Offset int
	Offset_size int
	Total int
	To interface{}
}

func (d *Data) InitBase(C *gin.Context){
	pa:= C.DefaultQuery("page", "1") //查询请求URL后面的参数，如果没有填写默认值
	pasi:= C.DefaultQuery("page_size", "5") //查询请求URL后面的参数，如果没有填写默认值

	d.Page, _ = strconv.Atoi(pa)
	d.Page_size, _ = strconv.Atoi(pasi)
	d.Offset=(d.Page-1)*d.Page_size
	d.Offset_size=d.Offset + d.Page_size
	//dd:=C.Query()
	return
}