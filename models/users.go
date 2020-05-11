package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//数据描述
type User struct {
	Id                  int    `json:"id" form:"id"`
	Uniacid 			int    `json:"uniacid" form:"uniacid"`
	Uid             	int    `json:"uid" form:"uid"`
	Openid          	string `json:"openid" form:"openid"`
	Realname        	string `json:"realname" form:"realname"`
	Mobile          	string `json:"mobile" form:"mobile"`
	Nickname        	string `json:"nickname" form:"nickname"`
	Nickname_wechat 	string `json:"nickname_wechat" form:"nickname_wechat"`
	Avatar       		string `json:"avatar" form:"avatar"`
	Avatar_wechat   	string `json:"avatar_wechat" form:"avatar_wechat"`
	Gender          	string `json:"gender" form:"gender"`
	Province        	string `json:"province" form:"province"`
	City            	string `json:"city" form:"city"`
	Area            	string `json:"area" form:"area"`
	Createtime      	int64  `json:"createtime" form:"createtime"`
	Agenttime       	int64  `json:"agenttime" form:"agenttime"`
	Status          	int8   `json:"status" form:"status"`
	Isagent         	int8   `json:"isagent" form:"isagent"`
	db *gorm.DB
}

func init() {
	Db=Db.Table(tablename("ewei_shop_member"))
}
func Table() User{
	return User{db:Db}
}
func Fields(s []string) User {
	return User{db:Db.Select(s)}
}

func Where(m map[string]interface{}) User {
	return User{db:Db.Where(m)}
}

func (u User) AddUsers(us *User) (id int64, err error) {
	err=u.db.Create(&us).Error
	if err != nil {
		return
	}
	id,err= strconv.ParseInt(strconv.Itoa(us.Id), 10, 64)
	if err != nil {
		return
	}
	return
}

func (u User) OneUsers() (users User)  {
	err:=u.db.First(&users).Error
	if err != nil {
		return
	}
	return
}

func Close() {
	if Db != nil {
		Db.Close()
	}
}
//
//// GetOrdersById retrieves Orders by Id. Returns error if
//// Id doesn't exist
//func GetOrdersById(id int) (v *Orders, err error) {
//	o := orm.NewOrm()
//	v = &Orders{Id: id}
//	if err = o.Read(v); err == nil {
//		return v, nil
//	}
//	return nil, err
//}
//
//// GetAllOrders retrieves all Orders matches certain condition. Returns empty list if
//// no records exist
//func GetAllOrders(query map[string]string, fields []string, sortby []string, order []string,
//	offset int64, limit int64) (ml []interface{}, err error) {
//	o := orm.NewOrm()
//	qs := o.QueryTable(new(Orders))
//	// query k=v
//	for k, v := range query {
//		// rewrite dot-notation to Object__Attribute
//		k = strings.Replace(k, ".", "__", -1)
//		if strings.Contains(k, "isnull") {
//			qs = qs.Filter(k, (v == "true" || v == "1"))
//		} else {
//			qs = qs.Filter(k, v)
//		}
//	}
//	// order by:
//	var sortFields []string
//	if len(sortby) != 0 {
//		if len(sortby) == len(order) {
//			// 1) for each sort field, there is an associated order
//			for i, v := range sortby {
//				orderby := ""
//				if order[i] == "desc" {
//					orderby = "-" + v
//				} else if order[i] == "asc" {
//					orderby = v
//				} else {
//					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//			qs = qs.OrderBy(sortFields...)
//		} else if len(sortby) != len(order) && len(order) == 1 {
//			// 2) there is exactly one order, all the sorted fields will be sorted by this order
//			for _, v := range sortby {
//				orderby := ""
//				if order[0] == "desc" {
//					orderby = "-" + v
//				} else if order[0] == "asc" {
//					orderby = v
//				} else {
//					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//		} else if len(sortby) != len(order) && len(order) != 1 {
//			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
//		}
//	} else {
//		if len(order) != 0 {
//			return nil, errors.New("Error: unused 'order' fields")
//		}
//	}
//
//	var l []Orders
//	qs = qs.OrderBy(sortFields...)
//	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
//		if len(fields) == 0 {
//			for _, v := range l {
//				ml = append(ml, v)
//			}
//		} else {
//			// trim unused fields
//			for _, v := range l {
//				m := make(map[string]interface{})
//				val := reflect.ValueOf(v)
//				for _, fname := range fields {
//					m[fname] = val.FieldByName(fname).Interface()
//				}
//				ml = append(ml, m)
//			}
//		}
//		return ml, nil
//	}
//	return nil, err
//}
//
//// UpdateOrders updates Orders by Id and returns error if
//// the record to be updated doesn't exist
//func UpdateOrdersById(m *Orders) (err error) {
//	o := orm.NewOrm()
//	v := Orders{Id: m.Id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Update(m); err == nil {
//			fmt.Println("Number of records updated in database:", num)
//		}
//	}
//	return
//}
//
//// DeleteOrders deletes Orders by Id and returns error if
//// the record to be deleted doesn't exist
//func DeleteOrders(id int) (err error) {
//	o := orm.NewOrm()
//	v := Orders{Id: id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Delete(&Orders{Id: id}); err == nil {
//			fmt.Println("Number of records deleted in database:", num)
//		}
//	}
//	return
//}
