package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	cf "myframe/conf"
)

var SqlDB *gorm.DB

func init() {
	var err error
	conf,err:=cf.GetArea("mysql")
	if err != nil {
		return
	}
	//"root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local
	SqlDB, err = gorm.Open(conf["DB_CONNECTION"], conf["DB_USERNAME"]+":"+conf["DB_PASSWORD"]+"@tcp("+conf["DB_HOST"]+":"+conf["DB_PORT"]+")/"+conf["DB_DATABASE"]+"?charset=utf8&parseTime=True&loc=Local")

	//SqlDB, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		return
	}
	SqlDB.SingularTable(true)
	return
}

func Close() {
	if SqlDB != nil {
		SqlDB.Close()
	}
}