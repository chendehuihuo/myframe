package conf

import (
	"github.com/Unknwon/goconfig"
)

func  GetArea(area string) ( map[string]string, error){
	var err error
	cfg, err := goconfig.LoadConfigFile("conf/app.conf")
	if err != nil{
		return nil,err
	}
	//Conf, err = cfg.GetValue(a.Area, a.Val)//读取单个值
	Conf ,err := cfg.GetSection(area)
	if err != nil{
		return nil,err
	}
	return Conf,nil
}