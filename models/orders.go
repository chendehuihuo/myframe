package models

//数据描述
type Orders struct {
	User string
	Password string
	Id                   int     `gorm:"column(id);auto" description:"订单ID"`
	OrdersSn             string  `gorm:"column(orders_sn);size(32)" description:"订单编号"`
	Type                 string  `gorm:"column(type);size(16);null" description:"积分，线上，线下，账期"`
	UserId               int     `gorm:"column(user_id)" description:"用户ID"`
	Configid             int     `gorm:"column(configid)" description:"下单时用户的级别id"`
	Isagree              int8    `gorm:"column(isagree)" description:"是否经过总账号同意，0通过；1未通过"`
	AgreeTime            int     `gorm:"column(agree_time)" description:"总账户同意时间"`
	AddressId            int     `gorm:"column(address_id)" description:"地址ID"`
	Consignee            string  `gorm:"column(consignee);size(64)" description:"收货人"`
	Country              string  `gorm:"column(country);size(256);null" description:"国家"`
	CountryCode          int     `gorm:"column(country_code);null"`
	Province             string  `gorm:"column(province);size(256);null" description:"省"`
	ProvinceCode         int     `gorm:"column(province_code);null"`
	CityCode             int     `gorm:"column(city_code);null"`
	City                 string  `gorm:"column(city);size(256);null" description:"市"`
	District             string  `gorm:"column(district);size(256);null" description:"区县编号"`
	DistrictCode         int     `gorm:"column(district_code);null"`
	Address              string  `gorm:"column(address);size(256);null" description:"地址"`
	Zipcode              string  `gorm:"column(zipcode);size(16);null" description:"邮编"`
	Telephone            string  `gorm:"column(telephone);size(16);null" description:"座机"`
	Mobile               string  `gorm:"column(mobile);size(16);null" description:"手机"`
	Remark               string  `gorm:"column(remark);size(1024);null" description:"留言"`
	ReceiptType          int8    `gorm:"column(receipt_type);null" description:"发票类型"`
	ReceiptId            int     `gorm:"column(receipt_id);null" description:"发票ID"`
	ReceiptTitle         string  `gorm:"column(receipt_title);size(256);null" description:"发票抬头"`
	GoodsAmount          float64 `gorm:"column(goods_amount);digits(10);decimals(2)" description:"商品总额"`
	GoodsNumber          int     `gorm:"column(goods_number);null"`
	ExpressAmount        float64 `gorm:"column(express_amount);null;digits(10);decimals(2)" description:"快递费"`
	TaxAmount            float64 `gorm:"column(tax_amount);null;digits(10);decimals(2)" description:"税"`
	OrdersAmount         float64 `gorm:"column(orders_amount);null;digits(10);decimals(2)" description:"订单总额"`
	Points               int     `gorm:"column(points)" description:"使用积分"`
	GivePoints           int     `gorm:"column(give_points);null" description:"赠送积分"`
	PointsAmount         float64 `gorm:"column(points_amount);digits(10);decimals(2)" description:"积分抵扣费用"`
	BonusSn              string  `gorm:"column(bonus_sn);size(64);null" description:"红包/优惠券码"`
	BonusType            int     `gorm:"column(bonus_type);null" description:"红包/优惠券类型"`
	BonusAmount          float64 `gorm:"column(bonus_amount);digits(10);decimals(2)" description:"优惠券/红包抵扣费用"`
	GiveRebate           float64 `gorm:"column(give_rebate);digits(10);decimals(2)" description:"赠与返点"`
	RebateAmount         float64 `gorm:"column(rebate_amount);digits(10);decimals(2)" description:"使用返利抵扣现金"`
	PaymentAmount        float64 `gorm:"column(payment_amount);null;digits(10);decimals(2)" description:"实际支付金额"`
	PaymentOrdersSn      string  `gorm:"column(payment_orders_sn);size(32);null" description:"当修改支付金额是生成支付订单号，用户微信支付"`
	PaymentTime          int     `gorm:"column(payment_time);null" description:"支付时间"`
	PaymentId            int8    `gorm:"column(payment_id);null" description:"支付方式"`
	PaymentName          string  `gorm:"column(payment_name);size(64);null" description:"支付方式名称"`
	PaymentStatus        int8    `gorm:"column(payment_status)" description:"支付状态 0 未支付 1 支付 2 支付失败"`
	ExpressId            int8    `gorm:"column(express_id);null" description:"快递方式"`
	ExpressName          string  `gorm:"column(express_name);size(255);null" description:"快递名字"`
	ExpressCom           string  `gorm:"column(express_com);size(32);null"`
	ExpressTime          int     `gorm:"column(express_time);null" description:"配送开始时间"`
	ExpressStatus        int     `gorm:"column(express_status)" description:"快递状态"`
	ExpressInvoiceNumber string  `gorm:"column(express_invoice_number);size(32);null" description:"发货单号"`
	ExpressFinishTime    int     `gorm:"column(express_finish_time)" description:"确认收货时间"`
	Status               int8    `gorm:"column(status)" description:"状态0正常；1取消；2失效 3删除"`
	StatusLast           int     `gorm:"column(status_last)" description:"记录上次status状态"`
	StatusChangeRemark   string  `gorm:"column(status_change_remark);size(1024);null" description:"状态变化说明，如取消订单"`
	IsComment            int8    `gorm:"column(is_comment);null" description:"是否已评价"`
	CreateTime           int     `gorm:"column(create_time);null" description:"创建时间"`
	UpdateTime           int     `gorm:"column(update_time)" description:"最后更新时间"`
	DeleteTime           int     `gorm:"column(delete_time)" description:"删除时间"`
}

func (t *Orders) TableName() string {
	return "orders"
}

//One
func (m *Orders) One() (one *Orders, err error) {
	//one = &Orders{}
	//err = crudOne(m, one)
	return
}
//
//// AddOrders insert a new Orders into database and returns
//// last inserted Id on success.
//func AddOrders(m *Orders) (id int64, err error) {
//	o := orm.NewOrm()
//	id, err = o.Insert(m)
//	return
//}
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
