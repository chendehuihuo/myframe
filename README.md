<p align="center">
    <a href="https://github.com/chendehuihuo" target="_blank">
        <img src="" height="100px">
    </a>
    <h1 align="center">LG</h1>
    <br>
</p>

LG基于Gin和Gorm结合的应用程序框架，最适合快速创建小项目。

该框架主要应用在接口方面使用



DIRECTORY STRUCTURE
-------------------

      apis/                 contains apis interface
      assets/               contains assets definition
      componnents/          contains console commands (apis)
      conf/                 contains application configurations
      database/             contains database connect
      models/               contains model classes
      public/               contains public resources
      routers/              contains Route definition

 安装
-------------------


    必须安装gin ： go get -u github.com/gin-gonic/gin

    必选解析conf文件：go get -u github.com/Unknwon/goconfig

    必选数据库二选一 ： go get -u github.com/go-sql-driver/mysql或go get -u github.com/jinzhu/gorm建议

    可选redis ： go get -u github.com/garyburd/redigo/redis

    可选微信支付 ： go get -u github.com/objcoding/wxpay

    当安装微信支付报错“golang.org/x/crypto/pkcs12”时，在使用
    go get -u golang.org/x/crypto/pkcs12 安装继续失败的话，在GOPATH mkdir golang.org/x
    
    目录golang.org/x 下 git clone  https://github.com/golang/crypto.git
    go install crypto
    之后再安装一遍  go get -u github.com/objcoding/wxpay
    
    可选生成二维码 ： go get -u github.com/skip2/go-qrcode

    可选支付宝支付 ： go get -u github.com/smartwalle/alipay
