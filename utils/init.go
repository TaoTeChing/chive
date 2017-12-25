package utils

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/liu2hai/chive/config"
	"github.com/liu2hai/chive/logs"

	"encoding/json"
)

func InitLogger(app string, level int) {
	host, _ := os.Hostname()
	year, month, day := time.Now().Date()
	datestr := fmt.Sprintf("%4d-%02d-%02d", year, month, day)
	fileName := fmt.Sprintf("%s/%s%d.%s.%s.log", config.T.LogPath, app, config.T.AppID, datestr, host)

	tag := &logs.LogCnfTag{
		Filename: fileName,
		MaxLines: 50000,
		Daily:    true, // Rotate daily
		MaxDays:  1,
		Rotate:   true,
		Level:    level,
		Perm:     "0666",
	}
	bs, _ := json.Marshal(tag)
	logs.SetLogger("file", string(bs))
	logs.SetLevel(level)
}

func InitCnf() {
	var appID *int = flag.Int("i", 1, "app id")
	var cnfPath *string = flag.String("c", "../lapf.cnf", "app config file")
	flag.Parse()

	config.T.AppID = *appID
	err := config.T.LoadConfig(*cnfPath)
	if err != nil {
		fmt.Println("can't load config file")
		panic(err)
	}
}

func InitOkexErrorMap(m map[int]string) {
	m[10000] = "必填参数为空"
	m[10001] = "参数错误"
	m[10002] = "验证失败"
	m[10003] = "该连接已经请求了其他用户的实时交易数据"
	m[10004] = "该连接没有请求此用户的实时交易数据"
	m[10005] = "api_key或者sign不合法"
	m[10008] = "非法参数"
	m[10009] = "订单不存在"
	m[10010] = "余额不足"
	m[10011] = "卖的数量小于BTC/LTC最小买卖额度"
	m[10012] = "当前网站暂时只支持btc_usd ltc_usd"
	m[10014] = "下单价格不得≤0或≥1000000"
	m[10015] = "暂不支持此channel订阅"
	m[10016] = "币数量不足"
	m[10017] = "WebSocket鉴权失败"
	m[10100] = "用户被冻结"
	m[10049] = "小额委托（<0.15BTC)的未成交委托数量不得大于50个"
	m[10216] = "非开放API"
	m[20001] = "用户不存在"
	m[20002] = "用户被冻结"
	m[20003] = "用户被爆仓冻结"
	m[20004] = "合约账户被冻结"
	m[20005] = "用户合约账户不存在"
	m[20006] = "必填参数为空"
	m[20007] = "参数错误"
	m[20008] = "合约账户余额为空"
	m[20009] = "虚拟合约状态错误"
	m[20010] = "合约风险率信息不存在"
	m[20011] = "开仓前保证金率超过90%"
	m[20012] = "开仓后保证金率超过90%"
	m[20013] = "暂无对手价"
	m[20014] = "系统错误"
	m[20015] = "订单信息不存在"
	m[20016] = "平仓数量是否大于同方向可用持仓数量"
	m[20017] = "非本人操作"
	m[20018] = "下单价格高于前一分钟的105%或低于95%"
	m[20019] = "该IP限制不能请求该资源"
	m[20020] = "密钥不存在"
	m[20021] = "指数信息不存在"
	m[20022] = "接口调用错误"
	m[20023] = "逐仓用户"
	m[20024] = "sign签名不匹配"
	m[20025] = "杠杆比率错误"
	m[20100] = "请求超时"
	m[20101] = "数据格式无效"
	m[20102] = "登录无效"
	m[20103] = "数据事件类型无效"
	m[20104] = "数据订阅类型无效"
	m[20107] = "JSON格式错误"
	m[20115] = "quote参数未匹配到"
	m[20116] = "参数不匹配"
	m[1002] = "交易金额大于余额"
	m[1003] = "交易金额小于最小交易值"
	m[1004] = "交易金额小于0"
	m[1007] = "没有交易市场信息"
	m[1008] = "没有最新行情信息"
	m[1009] = "没有订单"
	m[1010] = "撤销订单与原订单用户不一致"
	m[1011] = "没有查询到该用户"
	m[1013] = "没有订单类型"
	m[1014] = "没有登录"
	m[1015] = "没有获取到行情深度信息"
	m[1017] = "日期参数错误"
	m[1018] = "下单失败"
	m[1019] = "撤销订单失败"
	m[1024] = "币种不存在"
	m[1025] = "没有K线类型"
	m[1026] = "没有基准币数量"
	m[1027] = "参数不合法可能超出限制"
	m[1028] = "保留小数位失败"
	m[1029] = "正在准备中"
	m[1030] = "有融资融币无法进行交易"
	m[1031] = "转账余额不足"
	m[1032] = "该币种不能转账"
	m[1035] = "密码不合法"
	m[1036] = "谷歌验证码不合法"
	m[1037] = "谷歌验证码不正确"
	m[1038] = "谷歌验证码重复使用"
	m[1039] = "短信验证码输错限制"
	m[1040] = "短信验证码不合法"
	m[1041] = "短信验证码不正确"
	m[1042] = "谷歌验证码输错限制"
	m[1043] = "登陆密码不允许与交易密码一致"
	m[1044] = "原密码错误"
	m[1045] = "未设置二次验证"
	m[1046] = "原密码未输入"
	m[1048] = "用户被冻结"
	m[1201] = "账号零时删除"
	m[1202] = "账号不存在"
	m[1203] = "转账金额大于余额"
	m[1204] = "不同种币种不能转账"
	m[1205] = "账号不存在主从关系"
	m[1206] = "提现用户被冻结"
	m[1207] = "不支持转账"
	m[1208] = "没有该转账用户"
	m[1209] = "当前api不可用"
}
