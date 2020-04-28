package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/objcoding/wxpay"
	"github.com/skip2/go-qrcode"
	"github.com/smartwalle/alipay"
	"os"
	cf "myframe/conf"
)

func PayWx(c *gin.Context) {
	orders_sn := c.Request.FormValue("orders_sn")
	body := c.Request.FormValue("body")
	//"go get github.com/objcoding/wxpay"
	client := wxpay.NewClient(wxpay.NewAccount("wxdb3fe245d24b6f85", "1486973552", "071fbbe57d709c03445a55e2f18909ad", false)) // sandbox环境请传true

	// 统一下单

	params := make(wxpay.Params)

	params.SetString("body", body).
		SetString("out_trade_no", orders_sn).
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("trade_type", "NATIVE")

	p, err := client.UnifiedOrder(params)
	if err != nil {
		jsonData(c, gin.H{"err": err})
	}
	filename := "public/img/qrcode/qr.png"
	err = qrcode.WriteFile(p["code_url"], qrcode.Highest, 256, "qr.png")
	err = os.Rename("qr.png", filename)
	if err != nil {
		jsonData(c, gin.H{"err": err})
	}
	conf, _ := cf.GetArea("mysql")
	jsonData(c, gin.H{"url": conf["APP_URL"] + filename})
}

func PayAli(c *gin.Context) {
	var aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv1PR5Dz74h/PwtMb2fYKgLSP+u8+JnRcbmGYGUldUodmW9b57uqLlOPE+PPqfrIHgL387BLGA9B4SIYLvCE6gtJjMqDJfG6sOrpOgGz/7hw4C5MQp9MxSrUIuvkbj+A8uR+THFTyzt+P1Freyol7iN5JB7JuSnY5kwqxz91Q86bZ1FL99mRvJ5e7eYV/+wO2OVrvUFGO6CcbQehuAfE8jXLD18J3srJRxwIlmz5Mvrdz48mWWGwObFLgS4jtQWI99/O8rhcjU8X+WOABr43A81bMhA+SJboWWeibYbILeaG+xQZPw1hwXiLoaoR9OoaGOR5l4IaK/hMk2onaIzg/swIDAQAB" // 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
	var privateKey ="MIIEogIBAAKCAQEAv1PR5Dz74h/PwtMb2fYKgLSP+u8+JnRcbmGYGUldUodmW9b57uqLlOPE+PPqfrIHgL387BLGA9B4SIYLvCE6gtJjMqDJfG6sOrpOgGz/7hw4C5MQp9MxSrUIuvkbj+A8uR+THFTyzt+P1Freyol7iN5JB7JuSnY5kwqxz91Q86bZ1FL99mRvJ5e7eYV/+wO2OVrvUFGO6CcbQehuAfE8jXLD18J3srJRxwIlmz5Mvrdz48mWWGwObFLgS4jtQWI99/O8rhcjU8X+WOABr43A81bMhA+SJboWWeibYbILeaG+xQZPw1hwXiLoaoR9OoaGOR5l4IaK/hMk2onaIzg/swIDAQABAoIBAG4prhXTw1hs5qyP715tm5iF3s0micuMflgUG4abxn8pdRd2TYxyQhJln1mU4DU4/rbzLbkkCh3SIB/qBsG5BD6lMvbIu5rJAqQB+AdcO3R3jRGlCRcItD4GxNMbcfJCs9QHySWsB+5zrRYhiFxoEtg43SXwaC65qebvZniAPwc60JCmOLJiJOLLCsUwLpdo41gV4mlDmgl/pxRi7pN91JKMj/0WAUknGcgvtVp1mgEvenG/4vYRomnas0/0QNIQ/QWyNHw/FoO+oM7hV+yma1QeY1H8zXVQy5Pl/mbW8PjrqFSpSBBTCNH2F+27YTdXSpJn6OhwHMQso2ogE/xWzlkCgYEA7HIHFHKzZmW2dsU4OfLH6ZOpXEZ2itvVj640l5xNCiql6O0Oi4y2EcM0f+zY4cLJKXLyDSDQYWRxQPWcIptHbFBn0eDgZg+ZS9vM5+0VwOzNVfHAMH+d9A01/cQ1QM6Xu4PngllBU2tjMZET6dvsubMSCEfEZ8d9qD3K8mZxMj8CgYEAzyaQQpmNGML3Rbhp0qE0yUhQpVPbYwvB7zH4S++3bDRnCPFg44EZxm8hzsdLV9/lOsuERO9dzRid+D6K/ibnN9Uul9lFrEC2tdRn7lWWaeG93GmlIU6iIOW5VjTRuUgBrhj3sZFxygFV3KNm4addOUvGpIsDdfUkKQh7RDdVrY0CgYBMLsFROVPMXmFYC1t+cFcxTz52WM3BF+vLi5mub9HywKMMVbUmCYRwXLEBttJ6i0Q9cqfCyxvZBX/oKjrQe6+bmpyY5xfmSoheBMgUwjMMIkBWvLnr5nKCO+91gE6YPvhslqd1PETsF+GyCsibxRyVCLl6qYUC2ygFmUNWruMyCQKBgFbxBocNT0zfN00J01BV3IPGWCKriV6OpryvqlikqcU+UShNtq6Flr3UEspoltSw4dLFbCRKmoHkHRj9lvSIe27RiJau2DEtkAHMvjonrdOOrVL+/BF4oU9RW3I1BPjQofPiAsOCGdIAsXrDEgEPb28GwL56od82HhEiwCmnV8xZAoGADZYjfE9C9C2+m1/dakCnauZyvSOiiqO+N7eAi692SE6RLb4OTeGUXfqwBbt+sDqOpce8emqmc3hPtLGD55RYI2dOV4c0E77APoYhWTnAvxRPiYvpQeqQPxjxczCtxeHrqrqTkpWgB8XGjEBz9Ytfo/MkPBT1y48Oppqy7g6YFD4="
	var client,_ = alipay.New("2016080500172164", privateKey, false);

	var _ = client.LoadAliPayPublicKey(aliPublicKey)

	var p = alipay.TradeWapPay{}
	//var p = alipay.AliPayTradePagePay{}
	//var p = alipay.AliPayTradeAppPay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "QUICK_WAP_WAY"

	var url, _= client.TradeWapPay(p)
	//var url, err= client.TradePagePay(p)
	//if url == nil {
	//	jsonData(c,gin.H{"html":err})
	//}
	var payURL = url.String()
	//将html输出到浏览器

	jsonData(c, gin.H{"html": payURL})
}
