package apis

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	rd "myframe/assets"
	"myframe/models"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "users/login.html", gin.H{"title": "golang views", "header": "this is header", "footer": "this is footer"})
}

func Ajax_login(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	code := c.Request.FormValue("code")
	mid := c.Request.FormValue("mid")

	if mobile == "" || code == "" {
		jsonError(c, "缺少参数")
		return
	}

	var _, error = regexp.MatchString("/^1[3-9]\\d{9}$/", mobile)
	if error != nil {
		jsonError(c, "请填写正确的手机号格式")
		return
	}

	var iscode, _ = regexp.MatchString("\\d{5}", code)
	if !iscode {
		jsonError(c, "请填写正确的验证码格式")
		return
	}

	//var uid = 0
	//var mc = map[string]interface{}{};

	//$uid = mc_openid2uid($openid);
	//
	//$userinfo = array('openid' => $member['openid'], 'nickname' => $member['nickname'], 'headimgurl' => $member['avatar'], 'gender' => $member['gender'], 'province' => $member['province'], 'city' => $member['city']);

	var openid = "fake_" + MD5(mobile+strconv.FormatInt(time.Now().Unix(), 10));

	//$mc = array();
	//$mc['nickname'] = $userinfo['nickname'];
	//$mc['avatar'] = $userinfo['headimgurl'];
	//$mc['gender'] = $userinfo['sex'];
	//$mc['resideprovince'] = $userinfo['province'];
	//$mc['residecity'] = $userinfo['city'];
	//
	//} else {
	//// $uid = ;
	//}

	// 一个手机号对一个公众号唯一
	var has_exists=models.Where(map[string]interface{}{
		"mobile":mobile,
		"uniacid":2,
	}).OneUsers()
	// 邀请人存在
	var invite_info = models.User{}
	if mid != "" {
		invite_info=models.Where(map[string]interface{}{
			"id":mid,
			"uniacid":2,
		}).OneUsers()
	}

	conn := rd.Redis
	// 已经注册的会员
	if has_exists.Id != 0 {
		_, err := conn.Do("SET", "has_openid", has_exists.Openid)
		var rand = strconv.Itoa(rand.Intn(10000))
		_, err = conn.Do("SET", "has_rand", rand)
		_, err = conn.Do("SET", "mid", has_exists.Id)
		_, err = conn.Do("SET", "ws:login:user_token:"+has_exists.Openid, MD5(has_exists.Openid+"wsshop"+rand))
		if err != nil {
			jsonError(c, err.Error())
			return
		}
		jsonData(c, gin.H{"showInvite": false});
		return
	} else {
		//// 首次注册的会员
		 var u = models.User{
			Uniacid:         2,
			Uid:             0,
			Openid:          openid,
			Realname:        "",
			Mobile:          mobile,
			Nickname:        mobile,
			Nickname_wechat: "友品会员",
			Avatar:          "https://ws-shop.oss-cn-hangzhou.aliyuncs.com/images/3/2020/01/head.png",
			Avatar_wechat:   "",
			Gender:          "-1",
			Province:        "",
			City:            "",
			Area:            "",
			Createtime:      time.Now().Unix(),
			Agenttime:       time.Now().Unix(),
			Status:          0,
			Isagent:         0,
		}

		var userId, _ = models.Table().AddUsers(&u)
		if userId != 0 {
			_, err := conn.Do("SET", "has_openid", has_exists.Openid)
			var rand = strconv.Itoa(rand.Intn(10000))
			_, err = conn.Do("SET", "has_rand", rand)
			_, err = conn.Do("SET", "mid", has_exists.Id)
			_, err = conn.Do("SET", "ws:login:user_token:"+has_exists.Openid, MD5(has_exists.Openid+"wsshop"+rand))
			if err != nil {
				jsonError(c, err.Error())
				return
			}
			if mid != "" {
				// 通过邀请码进来的新用户
				jsonData(c, gin.H{"showInvite": true, "incode": mid, "isInvite": true, "invite_info": invite_info,"id":userId});
				return
			} else {
				// 不是通过邀请码进来的新用户
				jsonData(c, gin.H{"showInvite": true, "incode": "", "isInvite": false,"id":userId});
				return
			}
		} else {
			jsonError(c, "注册失败，请重试！");
			return
		}
	}
}
