package dataapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"util"
)

// 成语字典
func IdiomReq(word string) {

	var APPKEY = "6b7b03c4f06c7ba11cc036d74ee9dd07"

	//请求地址
	juheURL := "http://v.juhe.cn/chengyu/query"

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("word", word)    //填写需要查询的汉字，UTF8 urlencode编码
	param.Set("key", APPKEY)   //应用APPKEY(应用详细页查询)
	param.Set("dtype", "json") //返回数据的格式,xml或json，默认json

	//发送请求
	data, err := util.Get(juheURL, param)
	if err != nil {
		fmt.Errorf("请求失败,错误信息:\r\n%v", err)

	} else {
		var netReturn map[string]interface{}
		json.Unmarshal(data, &netReturn)
		if netReturn["error_code"].(float64) == 0 {
			fmt.Printf("接口返回result字段是:\r\n%v", netReturn["result"])
		}
	}
}

// 新华字典请求
